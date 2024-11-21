
#define LOAD_CONST 79

PyObject* _Py_HOT_FUNCTION
_PyEval_EvalFrameDefault(PyThreadState *tstate, 
                         _PyInterpreterFrame *frame, 
                         int throwflag)
{
    DISPATCH_GOTO();  // разворачивается в `goto dispatch_opcode`
dispatch_opcode:
    switch (opcode) {
        TARGET(LOAD_CONST):  // разворачивается в `case 79:`
        {   
            frame->instr_ptr = next_instr;
            next_instr += 1;
            _PyStackRef value = PyStackRef_FromPyObjectNew(
                GETITEM(FRAME_CO_CONSTS, oparg));
            // ...
        }
        // ...
    }

    opcode = next_instr->op.code;
    DISPATCH_GOTO();  // разворачивается в `goto dispatch_opcode;`

exit:
    // end of cycle: success or error
}
