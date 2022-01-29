package celeritas

const version = "0.0.1"
type Celeritas struct {
	AppName string
	Debug bool
	Version string

}
func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{rootPath: rootPath, folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}
func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, v := range p.folderNames {
		err := c.CreateDirIfNotExist(root +"/"+ v)
		if err != nil {
			return err
		}

	}
	return nil
}