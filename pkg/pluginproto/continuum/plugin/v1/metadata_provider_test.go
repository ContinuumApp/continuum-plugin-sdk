package pluginv1

import (
	"testing"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestMetadataItemDescriptor_IncludesReleaseDate(t *testing.T) {
	field := (&MetadataItem{}).ProtoReflect().Descriptor().Fields().ByName("release_date")
	if field == nil {
		t.Fatal("MetadataItem descriptor is missing release_date")
	}
}

func TestGetMetadataRequestDescriptor_IncludesContextFields(t *testing.T) {
	fields := (&GetMetadataRequest{}).ProtoReflect().Descriptor().Fields()
	for _, name := range []string{"provider_ids", "language", "file_path"} {
		if fields.ByName(protoreflect.Name(name)) == nil {
			t.Fatalf("GetMetadataRequest descriptor is missing %s", name)
		}
	}
}

func TestMetadataProviderRequestDescriptors_IncludeProviderContext(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		message   protoreflect.ProtoMessage
	}{
		{
			name:      "GetSeasonsRequest",
			fieldName: "provider_ids",
			message:   &GetSeasonsRequest{},
		},
		{
			name:      "GetEpisodesRequest",
			fieldName: "provider_ids",
			message:   &GetEpisodesRequest{},
		},
		{
			name:      "GetImagesRequest provider_ids",
			fieldName: "provider_ids",
			message:   &GetImagesRequest{},
		},
		{
			name:      "GetImagesRequest language",
			fieldName: "language",
			message:   &GetImagesRequest{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.message.ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name(tt.fieldName)) == nil {
				t.Fatalf("%s descriptor is missing %s", tt.name, tt.fieldName)
			}
		})
	}
}
