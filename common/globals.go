package common

const CODE_DUMPING = false
const JUMP_TRACING = false

// used for tracing

// CONTRACT_CODE contains the code of the contract being called so we can pass it to the Analysis stage
var CONTRACT_CODE = map[Hash][]byte{}

// CONTRACT_CODE_COUNT counts the amount of calls made to a specific contract
var CONTRACT_CODE_COUNT = map[Hash]uint{}
var CONTRACT_CODE_ALIAS = map[Address]Hash{}

// JMP_REG contains information about the jump commands that are executed
// using the following structure
// codeHash -> src -> dst -> callID -> counter
var JUMP_REG = map[Hash]map[uint32]map[uint32]map[int]uint{}
var JUMP_COUNT = map[Hash]uint{}
var JUMP_CALLS = map[Hash]map[int]struct{}{}

var CALLID = -1
var CALLID_COUNTER = 0
