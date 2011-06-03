// Copyright 2010 GoDCCP Authors. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package dccp

// abort() resets the connection with Reset Code 2, "Aborted"
func (c *Conn) abort() {
	c.Lock()
	c.socket.SetState(CLOSED)
	c.Unlock()
	c.inject(c.generateReset(ResetAborted))
	c.teardownUser()
	c.teardownWriteLoop()
}

// kill() kills the connection immediately and not gracefully
func (c *Conn) kill() {
	c.Lock()
	c.socket.SetState(CLOSED)
	c.Unlock()
	c.teardownUser()
	c.teardownWriteLoop()
}

func (c *Conn) teardownUser() {
	close(c.readApp)
	close(c.writeData)
}

func (c *Conn) teardownWriteLoop() {
	close(c.writeNonData)
}