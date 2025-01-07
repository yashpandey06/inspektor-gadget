// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package tracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type execsnoopEvent struct {
	MntnsId     uint64
	Timestamp   uint64
	Pid         uint32
	Tid         uint32
	Ppid        uint32
	Uid         uint32
	Gid         uint32
	Loginuid    uint32
	Sessionid   uint32
	Retval      int32
	ArgsCount   int32
	UpperLayer  bool
	PupperLayer bool
	_           [2]byte
	ArgsSize    uint32
	Comm        [16]uint8
	Pcomm       [16]uint8
	Args        [5120]uint8
	_           [4]byte
}

// loadExecsnoop returns the embedded CollectionSpec for execsnoop.
func loadExecsnoop() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_ExecsnoopBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load execsnoop: %w", err)
	}

	return spec, err
}

// loadExecsnoopObjects loads execsnoop and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*execsnoopObjects
//	*execsnoopPrograms
//	*execsnoopMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadExecsnoopObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadExecsnoop()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// execsnoopSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopSpecs struct {
	execsnoopProgramSpecs
	execsnoopMapSpecs
	execsnoopVariableSpecs
}

// execsnoopProgramSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopProgramSpecs struct {
	IgExecveE         *ebpf.ProgramSpec `ebpf:"ig_execve_e"`
	IgExecveX         *ebpf.ProgramSpec `ebpf:"ig_execve_x"`
	IgSchedExec       *ebpf.ProgramSpec `ebpf:"ig_sched_exec"`
	SecurityBprmCheck *ebpf.ProgramSpec `ebpf:"security_bprm_check"`
}

// execsnoopMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopMapSpecs struct {
	Events               *ebpf.MapSpec `ebpf:"events"`
	Execs                *ebpf.MapSpec `ebpf:"execs"`
	GadgetMntnsFilterMap *ebpf.MapSpec `ebpf:"gadget_mntns_filter_map"`
}

// execsnoopVariableSpecs contains global variables before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type execsnoopVariableSpecs struct {
	GadgetFilterByMntns *ebpf.VariableSpec `ebpf:"gadget_filter_by_mntns"`
	IgnoreFailed        *ebpf.VariableSpec `ebpf:"ignore_failed"`
	TargUid             *ebpf.VariableSpec `ebpf:"targ_uid"`
}

// execsnoopObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopObjects struct {
	execsnoopPrograms
	execsnoopMaps
	execsnoopVariables
}

func (o *execsnoopObjects) Close() error {
	return _ExecsnoopClose(
		&o.execsnoopPrograms,
		&o.execsnoopMaps,
	)
}

// execsnoopMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopMaps struct {
	Events               *ebpf.Map `ebpf:"events"`
	Execs                *ebpf.Map `ebpf:"execs"`
	GadgetMntnsFilterMap *ebpf.Map `ebpf:"gadget_mntns_filter_map"`
}

func (m *execsnoopMaps) Close() error {
	return _ExecsnoopClose(
		m.Events,
		m.Execs,
		m.GadgetMntnsFilterMap,
	)
}

// execsnoopVariables contains all global variables after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopVariables struct {
	GadgetFilterByMntns *ebpf.Variable `ebpf:"gadget_filter_by_mntns"`
	IgnoreFailed        *ebpf.Variable `ebpf:"ignore_failed"`
	TargUid             *ebpf.Variable `ebpf:"targ_uid"`
}

// execsnoopPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadExecsnoopObjects or ebpf.CollectionSpec.LoadAndAssign.
type execsnoopPrograms struct {
	IgExecveE         *ebpf.Program `ebpf:"ig_execve_e"`
	IgExecveX         *ebpf.Program `ebpf:"ig_execve_x"`
	IgSchedExec       *ebpf.Program `ebpf:"ig_sched_exec"`
	SecurityBprmCheck *ebpf.Program `ebpf:"security_bprm_check"`
}

func (p *execsnoopPrograms) Close() error {
	return _ExecsnoopClose(
		p.IgExecveE,
		p.IgExecveX,
		p.IgSchedExec,
		p.SecurityBprmCheck,
	)
}

func _ExecsnoopClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed execsnoop_x86_bpfel.o
var _ExecsnoopBytes []byte