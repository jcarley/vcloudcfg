package model

type Box struct {
	BoxName  string    `json:box-name`
	Versions []Version `json:versions`
}

func (b *Box) GetVersion(boxVersion string) *Version {
	return nil
}
