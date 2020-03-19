


// func systemstack(fn func())
TEXT runtime·systemstack(SB), NOSPLIT, $0-8
	MOVQ	fn+0(FP), DI	// DI = fn
	get_tls(CX)
	MOVQ	g(CX), AX	// AX = g
	MOVQ	g_m(AX), BX	// BX = m

	CMPQ	AX, m_gsignal(BX)
	JEQ	noswitch

	MOVQ	m_g0(BX), DX	// DX = g0
	CMPQ	AX, DX
	JEQ	noswitch

	CMPQ	AX, m_curg(BX)
	JNE	bad

	// switch stacks
	// save our state in g->sched. Pretend to
	// be systemstack_switch if the G stack is scanned.
	MOVQ	$runtime·systemstack_switch(SB), SI
	MOVQ	SI, (g_sched+gobuf_pc)(AX)
	MOVQ	SP, (g_sched+gobuf_sp)(AX)
	MOVQ	AX, (g_sched+gobuf_g)(AX)
	MOVQ	BP, (g_sched+gobuf_bp)(AX)

	// switch to g0
	MOVQ	DX, g(CX)
	MOVQ	(g_sched+gobuf_sp)(DX), BX
	// make it look like mstart called systemstack on g0, to stop traceback
	SUBQ	$8, BX
	MOVQ	$runtime·mstart(SB), DX
	MOVQ	DX, 0(BX)
	MOVQ	BX, SP

	// call target function
	MOVQ	DI, DX
	MOVQ	0(DI), DI
	CALL	DI

	// switch back to g
	get_tls(CX)
	MOVQ	g(CX), AX
	MOVQ	g_m(AX), BX
	MOVQ	m_curg(BX), AX
	MOVQ	AX, g(CX)
	MOVQ	(g_sched+gobuf_sp)(AX), SP
	MOVQ	$0, (g_sched+gobuf_sp)(AX)
	RET

noswitch:
	// already on m stack; tail call the function
	// Using a tail call here cleans up tracebacks since we won't stop
	// at an intermediate systemstack.
	MOVQ	DI, DX
	MOVQ	0(DI), DI
	JMP	DI

bad:
	// Bad: g is not gsignal, not g0, not curg. What is it?
	MOVQ	$runtime·badsystemstack(SB), AX
	CALL	AX
	INT	$3