package db1c7

import (
	mcdf "github.com/AlkBur/openmcdf"
	"sync"
	"web1c/utils"
)

const (
	metadata = "1Cv7.md"
	userdata = "users.usr"
	userDir = "UsrDef"
)

type DB struct {
	mu             sync.RWMutex
	dir string
	users []*User
	closed bool

	//OLE
	md *mcdf.CompoundFile
	usr *mcdf.CompoundFile

	//--------
	metadata *Metadata
}

func Open(dir string) (*DB, error) {
	db := &DB{
		dir: 	dir,
		closed: false,
	}

	//db.mu.Lock()
	//defer b.wg.Unlock()

	md, err := utils.FindFile(db.dir, metadata)
	if err != nil {
		return nil, err
	}
	if db.md, err = mcdf.Open(md); err != nil {
		return nil, err
	}

	if err = db.readMD(); err != nil {
		return nil, err
	}

	usr_Dir, err := utils.FindDirectory(db.dir, userDir)
	if err != nil {
		return db, nil
	}
	usr, err := utils.FindFile(usr_Dir, userdata)
	if err != nil {
		return nil, err
	}
	if db.usr, err = mcdf.Open(usr); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB)readMD() error {
	//GUIDData - содержит в себе информацию о наследовании файлов конфигурации
	//Main Metadata Stream - cодержит основную информацию о всех остальных элементах конфигурации
	//TagStream - все то, что находится в свойствах основного дерева метаданных на закладках "Автор" и "Заставка"

	db.mu.Lock()
	defer db.mu.Unlock()

	st, err := db.md.RootStorage().GetStorage("Metadata")
	if err != nil {
		return err
	}
	md, err := st.GetStream("Main MetaData Stream")
	if err != nil {
		return err
	}
	data, err := md.GetData()
	if err != nil {
		return err
	}

	//Возможно это описание версии формата стуркуры
	//data[0]=255
	//data[1]=158
	//data[2]=7
	data = data[3:]
	db.metadata, err = UnmarshalJSON(DecodeWindows1251(data))

	return err
}
