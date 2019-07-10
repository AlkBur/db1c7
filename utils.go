package db1c7

import (
	"github.com/mailru/easyjson/jlexer"
	"golang.org/x/text/encoding/charmap"
	"log"
	"strconv"
)

func DecodeWindows1251(ba []uint8) []uint8 {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.Bytes(ba)
	return out
}

func EncodeWindows1251(ba []uint8) []uint8 {
	enc := charmap.Windows1251.NewEncoder()
	out, _ := enc.Bytes(ba)
	return []uint8(out)
}

func UnmarshalJSON(data []byte) (*Metadata, error) {
	var result *Metadata

	log.Println(string(data))

	r := jlexer.Lexer{Data: data}
	easyjsonEnecodeMetada(&r, result)

	return result, r.Error()
}

func easyjsonEnecodeMetada(in *jlexer.Lexer, md *Metadata) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}

	md = new(Metadata)

	in.Delim('{')
	for in.IsDelim('{') {
		in.Delim('{')
		class := in.UnsafeString()
		switch class {
		case "MainDataContDef":
			ReadMainDataContDef(in, md)
		case "TaskItem":
			ReadTaskItem(in, md)
		case "GenJrnlFldDef":
			ReadGenJrnlFld(in, md)
		case "DocSelRefObj":
			
		default:
			log.Println(class)
		}
		in.Delim('}')
		in.WantComma()
	}

	in.Delim('}')

	if isTopLevel {
		in.Consumed()
	}
}

func ReadMainDataContDef(in *jlexer.Lexer, md *Metadata) {
	i := 0
	in.WantComma()
	for !in.IsDelim('}') {
		v := in.UnsafeString()
		switch i {
		case 0:
			md.ID, _ = strconv.Atoi(v)
		case 1:
			md.Version, _ = strconv.Atoi(v)
		case 2:
			md.Unknown, _ = strconv.Atoi(v)
		}
		in.WantComma()

		i++
	}
}

func ReadTaskItem(in *jlexer.Lexer, md *Metadata) {
	md.TaskItem = new(Task)
	in.WantComma()

	for in.IsDelim('{') {

		in.Delim('{')
		i := 0

		for !in.IsDelim('}') {
			switch i {
			case 0:
				md.TaskItem.ID, _ = strconv.Atoi(in.UnsafeString())
			case 1:
				md.TaskItem.Name = in.String()
			case 2:
				md.TaskItem.Comment = in.String()
			case 3:
				md.TaskItem.Description = in.String()
			case 4:
				md.TaskItem.Unknown1 = in.String()
			case 5:
				md.TaskItem.Unknown2, _ = strconv.Atoi(in.UnsafeString())
			case 6:
				md.TaskItem.Unknown3, _ = strconv.Atoi(in.UnsafeString())
			case 7:
				if in.UnsafeString() == "0" {
					md.TaskItem.AllowDirectDletion = false
				}else {
					md.TaskItem.AllowDirectDletion = true
				}
			case 8:
				md.TaskItem.Unknown4, _ = strconv.Atoi(in.UnsafeString())
			case 9:
				md.TaskItem.Unknown5, _ = strconv.Atoi(in.UnsafeString())
			}

			i++
			in.WantComma()
		}
		in.Delim('}')
		in.WantComma()
	}
}

func ReadGenJrnlFld(in *jlexer.Lexer, md *Metadata) {
	md.JournalFld = nil
}

