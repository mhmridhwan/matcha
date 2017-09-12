// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gomatcha.io/matcha/pb/view/button.proto

// This CPP symbol can be defined to use imports that match up to the framework
// imports needed when using CocoaPods.
#if !defined(GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS)
 #define GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS 0
#endif

#if GPB_USE_PROTOBUF_FRAMEWORK_IMPORTS
 #import <Protobuf/GPBProtocolBuffers_RuntimeSupport.h>
#else
 #import "GPBProtocolBuffers_RuntimeSupport.h"
#endif

 #import "gomatcha.io/matcha/pb/view/Button.pbobjc.h"
 #import "gomatcha.io/matcha/pb/Image.pbobjc.h"
// @@protoc_insertion_point(imports)

#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"

#pragma mark - MatchaViewPbButtonRoot

@implementation MatchaViewPbButtonRoot

// No extensions in the file and none of the imports (direct or indirect)
// defined extensions, so no need to generate +extensionRegistry.

@end

#pragma mark - MatchaViewPbButtonRoot_FileDescriptor

static GPBFileDescriptor *MatchaViewPbButtonRoot_FileDescriptor(void) {
  // This is called by +initialize so there is no need to worry
  // about thread safety of the singleton.
  static GPBFileDescriptor *descriptor = NULL;
  if (!descriptor) {
    GPB_DEBUG_CHECK_RUNTIME_VERSIONS();
    descriptor = [[GPBFileDescriptor alloc] initWithPackage:@"matcha.view"
                                                 objcPrefix:@"MatchaViewPb"
                                                     syntax:GPBFileSyntaxProto3];
  }
  return descriptor;
}

#pragma mark - MatchaViewPbButton

@implementation MatchaViewPbButton

@dynamic str;
@dynamic enabled;
@dynamic hasColor, color;

typedef struct MatchaViewPbButton__storage_ {
  uint32_t _has_storage_[1];
  NSString *str;
  MatchaPBColor *color;
} MatchaViewPbButton__storage_;

// This method is threadsafe because it is initially called
// in +initialize for each subclass.
+ (GPBDescriptor *)descriptor {
  static GPBDescriptor *descriptor = nil;
  if (!descriptor) {
    static GPBMessageFieldDescription fields[] = {
      {
        .name = "str",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPbButton_FieldNumber_Str,
        .hasIndex = 0,
        .offset = (uint32_t)offsetof(MatchaViewPbButton__storage_, str),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeString,
      },
      {
        .name = "enabled",
        .dataTypeSpecific.className = NULL,
        .number = MatchaViewPbButton_FieldNumber_Enabled,
        .hasIndex = 1,
        .offset = 2,  // Stored in _has_storage_ to save space.
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeBool,
      },
      {
        .name = "color",
        .dataTypeSpecific.className = GPBStringifySymbol(MatchaPBColor),
        .number = MatchaViewPbButton_FieldNumber_Color,
        .hasIndex = 3,
        .offset = (uint32_t)offsetof(MatchaViewPbButton__storage_, color),
        .flags = GPBFieldOptional,
        .dataType = GPBDataTypeMessage,
      },
    };
    GPBDescriptor *localDescriptor =
        [GPBDescriptor allocDescriptorForClass:[MatchaViewPbButton class]
                                     rootClass:[MatchaViewPbButtonRoot class]
                                          file:MatchaViewPbButtonRoot_FileDescriptor()
                                        fields:fields
                                    fieldCount:(uint32_t)(sizeof(fields) / sizeof(GPBMessageFieldDescription))
                                   storageSize:sizeof(MatchaViewPbButton__storage_)
                                         flags:GPBDescriptorInitializationFlag_None];
    NSAssert(descriptor == nil, @"Startup recursed!");
    descriptor = localDescriptor;
  }
  return descriptor;
}

@end


#pragma clang diagnostic pop

// @@protoc_insertion_point(global_scope)
