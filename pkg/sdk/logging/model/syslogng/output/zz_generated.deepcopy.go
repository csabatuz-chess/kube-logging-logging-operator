//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package output

import (
	"github.com/cisco-open/operator-tools/pkg/secret"
	"github.com/kube-logging/logging-operator/pkg/sdk/logging/model/syslogng/filter"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ALTS) DeepCopyInto(out *ALTS) {
	*out = *in
	if in.TargetServiceAccounts != nil {
		in, out := &in.TargetServiceAccounts, &out.TargetServiceAccounts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ALTS.
func (in *ALTS) DeepCopy() *ALTS {
	if in == nil {
		return nil
	}
	out := new(ALTS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	if in.ALTS != nil {
		in, out := &in.ALTS, &out.ALTS
		*out = new(ALTS)
		(*in).DeepCopyInto(*out)
	}
	if in.ADC != nil {
		in, out := &in.ADC, &out.ADC
		*out = new(ADC)
		**out = **in
	}
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(Insecure)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bulk) DeepCopyInto(out *Bulk) {
	*out = *in
	if in.Bulk != nil {
		in, out := &in.Bulk, &out.Bulk
		*out = new(bool)
		**out = **in
	}
	if in.BulkByPassValidation != nil {
		in, out := &in.BulkByPassValidation, &out.BulkByPassValidation
		*out = new(bool)
		**out = **in
	}
	if in.BulkUnordered != nil {
		in, out := &in.BulkUnordered, &out.BulkUnordered
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bulk.
func (in *Bulk) DeepCopy() *Bulk {
	if in == nil {
		return nil
	}
	out := new(Bulk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskBuffer) DeepCopyInto(out *DiskBuffer) {
	*out = *in
	if in.Compaction != nil {
		in, out := &in.Compaction, &out.Compaction
		*out = new(bool)
		**out = **in
	}
	if in.MemBufLength != nil {
		in, out := &in.MemBufLength, &out.MemBufLength
		*out = new(int64)
		**out = **in
	}
	if in.MemBufSize != nil {
		in, out := &in.MemBufSize, &out.MemBufSize
		*out = new(int64)
		**out = **in
	}
	if in.QOutSize != nil {
		in, out := &in.QOutSize, &out.QOutSize
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskBuffer.
func (in *DiskBuffer) DeepCopy() *DiskBuffer {
	if in == nil {
		return nil
	}
	out := new(DiskBuffer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchOutput) DeepCopyInto(out *ElasticsearchOutput) {
	*out = *in
	in.HTTPOutput.DeepCopyInto(&out.HTTPOutput)
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchOutput.
func (in *ElasticsearchOutput) DeepCopy() *ElasticsearchOutput {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileOutput) DeepCopyInto(out *FileOutput) {
	*out = *in
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileOutput.
func (in *FileOutput) DeepCopy() *FileOutput {
	if in == nil {
		return nil
	}
	out := new(FileOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPOutput) DeepCopyInto(out *HTTPOutput) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
	out.Batch = in.Batch
	in.Password.DeepCopyInto(&out.Password)
	if in.ResponseAction != nil {
		in, out := &in.ResponseAction, &out.ResponseAction
		*out = make(filter.RawArrowMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPOutput.
func (in *HTTPOutput) DeepCopy() *HTTPOutput {
	if in == nil {
		return nil
	}
	out := new(HTTPOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogScaleOutput) DeepCopyInto(out *LogScaleOutput) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogScaleOutput.
func (in *LogScaleOutput) DeepCopy() *LogScaleOutput {
	if in == nil {
		return nil
	}
	out := new(LogScaleOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loggly) DeepCopyInto(out *Loggly) {
	*out = *in
	if in.Token != nil {
		in, out := &in.Token, &out.Token
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	in.SyslogOutput.DeepCopyInto(&out.SyslogOutput)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loggly.
func (in *Loggly) DeepCopy() *Loggly {
	if in == nil {
		return nil
	}
	out := new(Loggly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LokiOutput) DeepCopyInto(out *LokiOutput) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(Auth)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(filter.ArrowMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LokiOutput.
func (in *LokiOutput) DeepCopy() *LokiOutput {
	if in == nil {
		return nil
	}
	out := new(LokiOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MQTT) DeepCopyInto(out *MQTT) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MQTT.
func (in *MQTT) DeepCopy() *MQTT {
	if in == nil {
		return nil
	}
	out := new(MQTT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDB) DeepCopyInto(out *MongoDB) {
	*out = *in
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
	if in.Uri != nil {
		in, out := &in.Uri, &out.Uri
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	out.ValuePairs = in.ValuePairs
	out.Batch = in.Batch
	in.Bulk.DeepCopyInto(&out.Bulk)
	out.WriteConcern = in.WriteConcern
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDB.
func (in *MongoDB) DeepCopy() *MongoDB {
	if in == nil {
		return nil
	}
	out := new(MongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RawString) DeepCopyInto(out *RawString) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RawString.
func (in *RawString) DeepCopy() *RawString {
	if in == nil {
		return nil
	}
	out := new(RawString)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisOutput) DeepCopyInto(out *RedisOutput) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.CommandAndArguments != nil {
		in, out := &in.CommandAndArguments, &out.CommandAndArguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.RenderedCommand.DeepCopyInto(&out.RenderedCommand)
	out.Batch = in.Batch
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisOutput.
func (in *RedisOutput) DeepCopy() *RedisOutput {
	if in == nil {
		return nil
	}
	out := new(RedisOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Output) DeepCopyInto(out *S3Output) {
	*out = *in
	if in.AccessKey != nil {
		in, out := &in.AccessKey, &out.AccessKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKey != nil {
		in, out := &in.SecretKey, &out.SecretKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectKey != nil {
		in, out := &in.ObjectKey, &out.ObjectKey
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	out.ObjectKeyTimestamp = in.ObjectKeyTimestamp
	out.Template = in.Template
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Output.
func (in *S3Output) DeepCopy() *S3Output {
	if in == nil {
		return nil
	}
	out := new(S3Output)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SplunkHECOutput) DeepCopyInto(out *SplunkHECOutput) {
	*out = *in
	in.HTTPOutput.DeepCopyInto(&out.HTTPOutput)
	in.Token.DeepCopyInto(&out.Token)
	if in.ExtraHeaders != nil {
		in, out := &in.ExtraHeaders, &out.ExtraHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraQueries != nil {
		in, out := &in.ExtraQueries, &out.ExtraQueries
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SplunkHECOutput.
func (in *SplunkHECOutput) DeepCopy() *SplunkHECOutput {
	if in == nil {
		return nil
	}
	out := new(SplunkHECOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StringList) DeepCopyInto(out *StringList) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StringList.
func (in *StringList) DeepCopy() *StringList {
	if in == nil {
		return nil
	}
	out := new(StringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SumologicHTTPOutput) DeepCopyInto(out *SumologicHTTPOutput) {
	*out = *in
	if in.Collector != nil {
		in, out := &in.Collector, &out.Collector
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SumologicHTTPOutput.
func (in *SumologicHTTPOutput) DeepCopy() *SumologicHTTPOutput {
	if in == nil {
		return nil
	}
	out := new(SumologicHTTPOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SumologicSyslogOutput) DeepCopyInto(out *SumologicSyslogOutput) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SumologicSyslogOutput.
func (in *SumologicSyslogOutput) DeepCopy() *SumologicSyslogOutput {
	if in == nil {
		return nil
	}
	out := new(SumologicSyslogOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyslogOutput) DeepCopyInto(out *SyslogOutput) {
	*out = *in
	if in.CloseOnInput != nil {
		in, out := &in.CloseOnInput, &out.CloseOnInput
		*out = new(bool)
		**out = **in
	}
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SoKeepalive != nil {
		in, out := &in.SoKeepalive, &out.SoKeepalive
		*out = new(bool)
		**out = **in
	}
	if in.TemplateEscape != nil {
		in, out := &in.TemplateEscape, &out.TemplateEscape
		*out = new(bool)
		**out = **in
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLS)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskBuffer != nil {
		in, out := &in.DiskBuffer, &out.DiskBuffer
		*out = new(DiskBuffer)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyslogOutput.
func (in *SyslogOutput) DeepCopy() *SyslogOutput {
	if in == nil {
		return nil
	}
	out := new(SyslogOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLS) DeepCopyInto(out *TLS) {
	*out = *in
	if in.CaDir != nil {
		in, out := &in.CaDir, &out.CaDir
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.CaFile != nil {
		in, out := &in.CaFile, &out.CaFile
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyFile != nil {
		in, out := &in.KeyFile, &out.KeyFile
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.CertFile != nil {
		in, out := &in.CertFile, &out.CertFile
		*out = new(secret.Secret)
		(*in).DeepCopyInto(*out)
	}
	if in.PeerVerify != nil {
		in, out := &in.PeerVerify, &out.PeerVerify
		*out = new(bool)
		**out = **in
	}
	if in.UseSystemCertStore != nil {
		in, out := &in.UseSystemCertStore, &out.UseSystemCertStore
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLS.
func (in *TLS) DeepCopy() *TLS {
	if in == nil {
		return nil
	}
	out := new(TLS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValuePairs) DeepCopyInto(out *ValuePairs) {
	*out = *in
	out.Scope = in.Scope
	out.Exclude = in.Exclude
	out.Key = in.Key
	out.Pair = in.Pair
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValuePairs.
func (in *ValuePairs) DeepCopy() *ValuePairs {
	if in == nil {
		return nil
	}
	out := new(ValuePairs)
	in.DeepCopyInto(out)
	return out
}
