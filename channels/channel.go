package channels

import "microAPro/models"

var MessageContextChannel = make(chan models.MessageContext, 100)
