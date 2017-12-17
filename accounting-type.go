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
