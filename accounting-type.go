package accounting

import "time"

type FrontPage struct {
	UserName string `json:"username"`
	LogOut   string `json:"logout"`
	Token    string `json:"token"`
}

type Input struct {
	Tanggal     time.Time `json:"tanggal"`
	NamaItem    string    `json:"nama"`
	Pemasukan   string    `json:"pemasukan"`
	Pengeluaran string    `json:"pengeluaran"`
	ServerMsg   string    `json:"servermsg"`
	Link        string    `json:"link"`
}

type ResponseJson struct {
	Data           interface{} `json:"data"`
	Script         string      `json:"script"`
	ModalScript    string      `json:"modal"`
	ScriptTambahan string      `json:"tambahan"`
}

type Kursor struct {
	Point string `json:"point"`
	Link  string `json:"link"`
}

type MaksPengeluaran struct {
	Tanggal         time.Time `json:"tanggal"`
	MaksPengeluaran int       `json:"makspengeluaran"`
}

type PengeluaranPage struct {
	TotalPengeluaran  int      `json:"total"`
	DaftarPengeluaran []Input  `json:"daftar"`
	DaftarKursor      []Kursor `json:"kursor"`
}

type PemasukanPage struct {
	PemasukanBulanLalu int      `json:"bulanlalu"`
	DaftarPemasukan    []Input  `json:"pemasukan"`
	TotalPemasukan     int      `json:"total"`
	PemasukanBulanIni  int      `json:"bulanini"`
	TotalTabungan      int      `json:"totaltabungan"`
	DaftarKursor       []Kursor `json:"kursor"`
}

type ResumePerbulan struct {
	Tanggal          time.Time `json:"tanggal"`
	Pengeluaran      int       `json:"pengeluaran"`
	Pemasukan        int       `json:"pemasukan"`
	TotalPengeluaran int       `json:"totalpeng"`
	TotalPemasukan   int       `json:"totalpem"`
}

type TukarJaga struct {
	TanggalInput      time.Time `json:"tglinput"`
	Nama              string    `json:"nama"`
	JagaHutang        string    `json:"jagahutang"`
	TanggalJagaHutang time.Time `json:"tgljagahutang"`
	JagaBayar         string    `json:"jagabayar"`
	TanggalJagaBayar  time.Time `json:"tgljagabayar"`
	Link              string    `json:"link"`
	Status            string    `json:"status"`
}

type CatchDataJson struct {
	Data1 string `json:"data01"`
	Data2 string `json:"data02"`
	Data3 string `json:"data03"`
	Data4 string `json:"data04"`
	Data5 string `json:"data05"`
	Data6 string `json:"data06"`
}
