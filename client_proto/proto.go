// Code generated by gen_proto.sh.
// DO NOT EDIT!
package client_proto

import "landlords/misc/packet"
import "fmt"
import "strings"

type S_null_struct struct {
}

func (p S_null_struct) Pack(w *packet.Packet) {
}

func PKT_null_struct(reader *packet.Packet) (tbl S_null_struct, err error) {

	return
}

type S_byte_id struct {
	F_id uint8
}

func (p S_byte_id) Pack(w *packet.Packet) {
	w.WriteByte(p.F_id)
}

func PKT_byte_id(reader *packet.Packet) (tbl S_byte_id, err error) {
	tbl.F_id, err = reader.ReadByte()
	checkErr(err)

	return
}

type S_auto_id struct {
	F_id int32
}

func (p S_auto_id) Pack(w *packet.Packet) {
	w.WriteS32(p.F_id)
}

func PKT_auto_id(reader *packet.Packet) (tbl S_auto_id, err error) {
	tbl.F_id, err = reader.ReadS32()
	checkErr(err)

	return
}

type S_entity_id struct {
	F_id string
}

func (p S_entity_id) Pack(w *packet.Packet) {
	w.WriteString(p.F_id)
}

func PKT_entity_id(reader *packet.Packet) (tbl S_entity_id, err error) {
	tbl.F_id, err = reader.ReadString()
	checkErr(err)
	tbl.F_id = strings.TrimSpace(tbl.F_id)

	return
}

type S_item_id struct {
	F_id uint32
}

func (p S_item_id) Pack(w *packet.Packet) {
	w.WriteU32(p.F_id)
}

func PKT_item_id(reader *packet.Packet) (tbl S_item_id, err error) {
	tbl.F_id, err = reader.ReadU32()
	checkErr(err)

	return
}

type S_player_card struct {
	F_hole_cards []string
	F_roomId     string
	F_players    []S_player
}

func (p S_player_card) Pack(w *packet.Packet) {
	w.WriteU16(uint16(len(p.F_hole_cards)))
	for k := range p.F_hole_cards {
		w.WriteString(p.F_hole_cards[k])
	}
	w.WriteString(p.F_roomId)
	w.WriteU16(uint16(len(p.F_players)))
	for k := range p.F_players {
		p.F_players[k].Pack(w)
	}
}

func PKT_player_card(reader *packet.Packet) (tbl S_player_card, err error) {
	{
		narr, err := reader.ReadU16()
		checkErr(err)
		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_hole_cards = append(tbl.F_hole_cards, v)
			checkErr(err)
		}
	}
	tbl.F_roomId, err = reader.ReadString()
	checkErr(err)
	tbl.F_roomId = strings.TrimSpace(tbl.F_roomId)
	{
		narr, err := reader.ReadU16()
		checkErr(err)
		tbl.F_players = make([]S_player, narr)
		for i := 0; i < int(narr); i++ {
			tbl.F_players[i], err = PKT_player(reader)
			checkErr(err)
		}
	}

	return
}

type S_player struct {
	F_id    string
	F_cards []string
}

func (p S_player) Pack(w *packet.Packet) {
	w.WriteString(p.F_id)
	w.WriteU16(uint16(len(p.F_cards)))
	for k := range p.F_cards {
		w.WriteString(p.F_cards[k])
	}
}

func PKT_player(reader *packet.Packet) (tbl S_player, err error) {
	tbl.F_id, err = reader.ReadString()
	checkErr(err)
	tbl.F_id = strings.TrimSpace(tbl.F_id)
	{
		narr, err := reader.ReadU16()
		checkErr(err)
		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_cards = append(tbl.F_cards, v)
			checkErr(err)
		}
	}

	return
}

type S_player_outof_card struct {
	F_roomId string
	F_cards  []string
}

func (p S_player_outof_card) Pack(w *packet.Packet) {
	w.WriteString(p.F_roomId)
	w.WriteU16(uint16(len(p.F_cards)))
	for k := range p.F_cards {
		w.WriteString(p.F_cards[k])
	}
}

func PKT_player_outof_card(reader *packet.Packet) (tbl S_player_outof_card, err error) {
	tbl.F_roomId, err = reader.ReadString()
	checkErr(err)
	tbl.F_roomId = strings.TrimSpace(tbl.F_roomId)
	{
		narr, err := reader.ReadU16()
		checkErr(err)
		for i := 0; i < int(narr); i++ {
			v, err := reader.ReadString()
			tbl.F_cards = append(tbl.F_cards, v)
			checkErr(err)
		}
	}

	return
}

func checkErr(err error) {
	if err != nil {
		panic("error occured in protocol module")
	}
}

func checkMax(path string, v float64, max float64) {
	if v > max {
		s := fmt.Sprintf("error range in %s, v=%g, max=%g", path, v, max)
		panic(s)
	}
}

func checkMin(path string, v float64, min float64) {
	if v < min {
		s := fmt.Sprintf("error range in %s, v=%g, min=%g", path, v, min)
		panic(s)
	}
}
