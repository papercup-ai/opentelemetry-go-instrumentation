// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package server

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfGrpcRequestT struct {
	StartTime  uint64
	EndTime    uint64
	Sc         bpfSpanContext
	Psc        bpfSpanContext
	Method     [100]int8
	StatusCode uint32
	LocalAddr  struct {
		Ip   [16]uint8
		Port uint32
	}
	_ [4]byte
}

type bpfSliceArrayBuff struct{ Buff [1024]uint8 }

type bpfSpanContext struct {
	TraceID    [16]uint8
	SpanID     [8]uint8
	TraceFlags uint8
	Padding    [7]uint8
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	UprobeHttp2ServerWriteStatus    *ebpf.ProgramSpec `ebpf:"uprobe_http2Server_WriteStatus"`
	UprobeHttp2ServerOperateHeader  *ebpf.ProgramSpec `ebpf:"uprobe_http2Server_operateHeader"`
	UprobeServerHandleStream        *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturns *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream_Returns"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	AllocMap              *ebpf.MapSpec `ebpf:"alloc_map"`
	Events                *ebpf.MapSpec `ebpf:"events"`
	GoContextToSc         *ebpf.MapSpec `ebpf:"go_context_to_sc"`
	GrpcEvents            *ebpf.MapSpec `ebpf:"grpc_events"`
	GrpcStorageMap        *ebpf.MapSpec `ebpf:"grpc_storage_map"`
	ProbeActiveSamplerMap *ebpf.MapSpec `ebpf:"probe_active_sampler_map"`
	SamplersConfigMap     *ebpf.MapSpec `ebpf:"samplers_config_map"`
	SliceArrayBuffMap     *ebpf.MapSpec `ebpf:"slice_array_buff_map"`
	StreamidToGrpcEvents  *ebpf.MapSpec `ebpf:"streamid_to_grpc_events"`
	TrackedSpansBySc      *ebpf.MapSpec `ebpf:"tracked_spans_by_sc"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	AllocMap              *ebpf.Map `ebpf:"alloc_map"`
	Events                *ebpf.Map `ebpf:"events"`
	GoContextToSc         *ebpf.Map `ebpf:"go_context_to_sc"`
	GrpcEvents            *ebpf.Map `ebpf:"grpc_events"`
	GrpcStorageMap        *ebpf.Map `ebpf:"grpc_storage_map"`
	ProbeActiveSamplerMap *ebpf.Map `ebpf:"probe_active_sampler_map"`
	SamplersConfigMap     *ebpf.Map `ebpf:"samplers_config_map"`
	SliceArrayBuffMap     *ebpf.Map `ebpf:"slice_array_buff_map"`
	StreamidToGrpcEvents  *ebpf.Map `ebpf:"streamid_to_grpc_events"`
	TrackedSpansBySc      *ebpf.Map `ebpf:"tracked_spans_by_sc"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.AllocMap,
		m.Events,
		m.GoContextToSc,
		m.GrpcEvents,
		m.GrpcStorageMap,
		m.ProbeActiveSamplerMap,
		m.SamplersConfigMap,
		m.SliceArrayBuffMap,
		m.StreamidToGrpcEvents,
		m.TrackedSpansBySc,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	UprobeHttp2ServerWriteStatus    *ebpf.Program `ebpf:"uprobe_http2Server_WriteStatus"`
	UprobeHttp2ServerOperateHeader  *ebpf.Program `ebpf:"uprobe_http2Server_operateHeader"`
	UprobeServerHandleStream        *ebpf.Program `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturns *ebpf.Program `ebpf:"uprobe_server_handleStream_Returns"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.UprobeHttp2ServerWriteStatus,
		p.UprobeHttp2ServerOperateHeader,
		p.UprobeServerHandleStream,
		p.UprobeServerHandleStreamReturns,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_x86_bpfel.o
var _BpfBytes []byte
