package bridge

import "testing"

func TestCommonMessageSendMessage(t *testing.T) {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("Doing nucleic acid?", "xxx")
	// log：Send Doing nucleic acid? to xxx via SMS
	t.Log("DONE")
}

func TestUrgencyMessageSendMessage(t *testing.T) {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("Doing nucleic acid?", "yyy")
	// log：Send [Urgency] Doing nucleic acid? to yyy via Email
	t.Log("DONE")
}
