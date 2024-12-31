package controllers

import "sync"

var lockMagicController = &sync.Mutex{}
