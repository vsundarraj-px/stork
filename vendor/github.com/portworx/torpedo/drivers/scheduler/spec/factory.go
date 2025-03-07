package spec

import (
	"fmt"
	"github.com/portworx/torpedo/pkg/log"
	"io/ioutil"
	"path"

	"github.com/portworx/torpedo/pkg/errors"
)

// Factory is an application spec factory
type Factory struct {
	specDir    string
	specParser Parser
}

var appSpecFactory = make(map[string]*AppSpec)

// register registers a new spec with the factory
func (f *Factory) register(id string, app *AppSpec) {
	if _, ok := appSpecFactory[id]; !ok {
		log.Debugf("Registering new app: %v", id)
	} else {
		log.Debugf("Substitute with new app: %v", id)
	}
	// NOTE: In case of spec rescan we need to substitute old app with another one
	appSpecFactory[id] = app
}

// Get returns a registered application
func (f *Factory) Get(id string) (*AppSpec, error) {
	if d, ok := appSpecFactory[id]; ok && d.Enabled {
		if copy := d.DeepCopy(); copy != nil {
			return d.DeepCopy(), nil
		}
		return nil, fmt.Errorf("error creating copy of app: %v", d)
	}

	return nil, &errors.ErrNotFound{
		ID:   id,
		Type: "AppSpec",
	}
}

// GetAll returns all registered enabled applications
func (f *Factory) GetAll() []*AppSpec {
	var specs []*AppSpec
	for _, val := range appSpecFactory {
		if val.Enabled {
			valCopy := val.DeepCopy()
			if valCopy != nil {
				specs = append(specs, valCopy)
			}
		}
	}

	return specs
}

// NewFactory creates a new spec factory
func NewFactory(specDir, storageProvisioner string, parser Parser) (*Factory, error) {
	f := &Factory{
		specDir:    specDir,
		specParser: parser,
	}

	appDirList, err := ioutil.ReadDir(f.specDir)
	if err != nil {
		return nil, err
	}

	for _, file := range appDirList {
		if file.IsDir() {

			specID := file.Name()

			specToParse := path.Join(f.specDir, specID)
			log.Debugf("Parsing: %v...", path.Join(f.specDir, specID))
			log.Debugf("Storage provisioner %s", storageProvisioner)
			specs, err := f.specParser.ParseSpecs(specToParse, storageProvisioner)
			if err != nil {
				return nil, err
			}

			if len(specs) == 0 {
				continue
			}

			// Register the spec
			f.register(specID, &AppSpec{
				Key:      specID,
				SpecList: specs,
				Enabled:  true,
			})
		}
	}

	if apps := f.GetAll(); len(apps) == 0 {
		return nil, fmt.Errorf("found 0 supported applications in given specDir: %v", specDir)
	}

	return f, nil
}
