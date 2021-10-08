package factory

import "golang.frontdoorhome.com/personal-project/tiny-url/internal/errors"

// GenericFactory Factory contains instances
type GenericFactory struct {
	factoryMap map[string]interface{}
}

// InitFactory Initialize the factory
func InitFactory() *GenericFactory {
	return &GenericFactory{factoryMap: make(map[string]interface{})}
}

// Register Register an object with the factory
func (factory *GenericFactory) Register(identifier string, object interface{}) (err error) {
	if identifier == "" {
		//logging.Log.Error(err)
		return &errors.EmptyObjectIdentifierError
	}

	_, objectExists := factory.factoryMap[identifier]
	if objectExists {
		//logging.Log.WithField("Tenant ID", tenantID).WithField("Factory identifier", identifier).Error(err)
		return &errors.ObjectAlreadyExistsInFactoryError
	}
	factory.factoryMap[identifier] = object

	return nil
}

// Get Fetch the object for the given tenant and identifier
func (factory *GenericFactory) Get(identifier string) (object interface{}, err error) {

	if identifier == "" {
		//	logging.Log.Error(err)
		return nil, &errors.EmptyObjectIdentifierError
	}

	object, objectExists := factory.factoryMap[identifier]
	if objectExists {
		return object, nil
	}
	//logging.Log.WithField("Tenant ID", tenantID).WithField("Factory identifier", identifier).Error(err)
	return nil, &errors.InvalidObjectIdentifierError

}
