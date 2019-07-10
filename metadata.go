package db1c7

type Metadata struct {
	ID      int
	Version int
	Unknown int
	//------------
	TaskItem        *Task  //Задача (пароль на конфигурацию и т.п)
	GenJrnlFldDef   struct{} //Общие реквизиты документов
	DocSelRefObj    struct{} //Графы отбора документов
	DocNumDef       struct{} //Нумераторы документов
	Consts          []*Const //Константы
	SbCnts          []*SbCnt
	Registers       []*Register
	Documents       []*Document //Документы. Тут все, кроме модуля проведения документа.
	Journalisters   []*Journalister
	EnumList        []*Enum   //Перечисления
	ReportList      []*Report //Отчеты
	CJ              interface{}
	Calendars       []*Calendar
	Algorithms      []*Algorithm
	RecalcRules     []*RecalcRule
	CalcVars        []*CalcVar
	Groups          []*Group
	DocumentStreams []*DocumentStream
	Buh             interface{}
	JournalFld      interface{}
	CRC             string //Контрольная сумма
}

type Task struct {
	ID          int
	Name        string
	Comment     string
	Description string
	Unknown1 		string   //возможно пароль
	Unknown2 		int   //возможно язык
	Unknown3 		int
	AllowDirectDletion bool //разрешить непосредственное удаление
	Unknown4 		int
	Unknown5 		int
}

type Const struct {
	ID int
	Name        string
	Comment     string
	Description string
	Type string
}

type SbCnt struct {
	ID int
	Name        string
	Comment     string
	Description string
}

type Register struct {

}

type Document struct {

}

type Journalister struct {

}

type Calendar struct {

}

type Enum struct {

}

type Report struct {

}

type Algorithm struct {

}

type RecalcRule struct {

}

type CalcVar struct {

}

type Group struct {

}

type DocumentStream struct {

}