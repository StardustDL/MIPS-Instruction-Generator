package exec

import (
    "../../instruction"
    "../cpu"
    "../memory"
)

func retrieveCode() uint32 {
    return memory.Read(cpu.PC,4) >> 6 & 0xfffff
}

func syscall(it rinstr) {
    if IsDebug{
        println("system call",cpu.GetGPR(instruction.GPR_V0))
    }
    doSystemCall(cpu.GetGPR(instruction.GPR_V0))
}

func _break(it rinstr){
    doBreak(retrieveCode())
}
