package channels

import "microAPro/models"

var MessageContextChannel = make(chan models.MessageContext, 100)

var AIChannel = make(chan AIAsk, 100)
