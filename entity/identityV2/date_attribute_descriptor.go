package identityV2

import "github.com/bloock/bloock-sdk-go/v2/internal/bridge/proto"

type DateAttributeDescriptor struct {
	AttributeDescriptor
}

func NewDateAttributeDescriptor(name string, id string, description string, required bool) DateAttributeDescriptor {
	a := AttributeDescriptor{
		DisplayName: name,
		Id:          id,
		Description: description,
		Required:    required,
	}
	return DateAttributeDescriptor{
		AttributeDescriptor: a,
	}
}

func NewDateAttributeDescriptorFromProto(s *proto.DateAttributeDefinitionV2) DateAttributeDescriptor {
	if s == nil {
		return DateAttributeDescriptor{}
	}
	a := AttributeDescriptor{
		DisplayName: s.GetId(),
		Id:          s.Id,
		Description: s.Description,
		Required:    s.Required,
	}
	return DateAttributeDescriptor{
		AttributeDescriptor: a,
	}
}

func (s DateAttributeDescriptor) ToProto() *proto.DateAttributeDefinitionV2 {
	return &proto.DateAttributeDefinitionV2{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		Description: s.Description,
		Required:    s.Required,
	}
}