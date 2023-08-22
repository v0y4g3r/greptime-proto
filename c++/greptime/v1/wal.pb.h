// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: greptime/v1/wal.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_greptime_2fv1_2fwal_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_greptime_2fv1_2fwal_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3021000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3021006 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
#include "greptime/v1/row.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_greptime_2fv1_2fwal_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_greptime_2fv1_2fwal_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_greptime_2fv1_2fwal_2eproto;
namespace greptime {
namespace v1 {
class BytesWrapper;
struct BytesWrapperDefaultTypeInternal;
extern BytesWrapperDefaultTypeInternal _BytesWrapper_default_instance_;
class Mutation;
struct MutationDefaultTypeInternal;
extern MutationDefaultTypeInternal _Mutation_default_instance_;
class WalEntry;
struct WalEntryDefaultTypeInternal;
extern WalEntryDefaultTypeInternal _WalEntry_default_instance_;
}  // namespace v1
}  // namespace greptime
PROTOBUF_NAMESPACE_OPEN
template<> ::greptime::v1::BytesWrapper* Arena::CreateMaybeMessage<::greptime::v1::BytesWrapper>(Arena*);
template<> ::greptime::v1::Mutation* Arena::CreateMaybeMessage<::greptime::v1::Mutation>(Arena*);
template<> ::greptime::v1::WalEntry* Arena::CreateMaybeMessage<::greptime::v1::WalEntry>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace greptime {
namespace v1 {

enum OpType : int {
  DELETE = 0,
  PUT = 1,
  OpType_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::min(),
  OpType_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::max()
};
bool OpType_IsValid(int value);
constexpr OpType OpType_MIN = DELETE;
constexpr OpType OpType_MAX = PUT;
constexpr int OpType_ARRAYSIZE = OpType_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* OpType_descriptor();
template<typename T>
inline const std::string& OpType_Name(T enum_t_value) {
  static_assert(::std::is_same<T, OpType>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function OpType_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    OpType_descriptor(), enum_t_value);
}
inline bool OpType_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, OpType* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<OpType>(
    OpType_descriptor(), name, value);
}
// ===================================================================

class Mutation final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:greptime.v1.Mutation) */ {
 public:
  inline Mutation() : Mutation(nullptr) {}
  ~Mutation() override;
  explicit PROTOBUF_CONSTEXPR Mutation(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Mutation(const Mutation& from);
  Mutation(Mutation&& from) noexcept
    : Mutation() {
    *this = ::std::move(from);
  }

  inline Mutation& operator=(const Mutation& from) {
    CopyFrom(from);
    return *this;
  }
  inline Mutation& operator=(Mutation&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const Mutation& default_instance() {
    return *internal_default_instance();
  }
  static inline const Mutation* internal_default_instance() {
    return reinterpret_cast<const Mutation*>(
               &_Mutation_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(Mutation& a, Mutation& b) {
    a.Swap(&b);
  }
  inline void Swap(Mutation* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Mutation* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Mutation* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Mutation>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Mutation& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const Mutation& from) {
    Mutation::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::PROTOBUF_NAMESPACE_ID::Arena* arena, bool is_message_owned);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Mutation* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "greptime.v1.Mutation";
  }
  protected:
  explicit Mutation(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kRowsFieldNumber = 3,
    kSequenceFieldNumber = 2,
    kOpTypeFieldNumber = 1,
  };
  // .greptime.v1.Rows rows = 3;
  bool has_rows() const;
  private:
  bool _internal_has_rows() const;
  public:
  void clear_rows();
  const ::greptime::v1::Rows& rows() const;
  PROTOBUF_NODISCARD ::greptime::v1::Rows* release_rows();
  ::greptime::v1::Rows* mutable_rows();
  void set_allocated_rows(::greptime::v1::Rows* rows);
  private:
  const ::greptime::v1::Rows& _internal_rows() const;
  ::greptime::v1::Rows* _internal_mutable_rows();
  public:
  void unsafe_arena_set_allocated_rows(
      ::greptime::v1::Rows* rows);
  ::greptime::v1::Rows* unsafe_arena_release_rows();

  // uint64 sequence = 2;
  void clear_sequence();
  uint64_t sequence() const;
  void set_sequence(uint64_t value);
  private:
  uint64_t _internal_sequence() const;
  void _internal_set_sequence(uint64_t value);
  public:

  // .greptime.v1.OpType op_type = 1;
  void clear_op_type();
  ::greptime::v1::OpType op_type() const;
  void set_op_type(::greptime::v1::OpType value);
  private:
  ::greptime::v1::OpType _internal_op_type() const;
  void _internal_set_op_type(::greptime::v1::OpType value);
  public:

  // @@protoc_insertion_point(class_scope:greptime.v1.Mutation)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::greptime::v1::Rows* rows_;
    uint64_t sequence_;
    int op_type_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_greptime_2fv1_2fwal_2eproto;
};
// -------------------------------------------------------------------

class WalEntry final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:greptime.v1.WalEntry) */ {
 public:
  inline WalEntry() : WalEntry(nullptr) {}
  ~WalEntry() override;
  explicit PROTOBUF_CONSTEXPR WalEntry(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  WalEntry(const WalEntry& from);
  WalEntry(WalEntry&& from) noexcept
    : WalEntry() {
    *this = ::std::move(from);
  }

  inline WalEntry& operator=(const WalEntry& from) {
    CopyFrom(from);
    return *this;
  }
  inline WalEntry& operator=(WalEntry&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const WalEntry& default_instance() {
    return *internal_default_instance();
  }
  static inline const WalEntry* internal_default_instance() {
    return reinterpret_cast<const WalEntry*>(
               &_WalEntry_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(WalEntry& a, WalEntry& b) {
    a.Swap(&b);
  }
  inline void Swap(WalEntry* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(WalEntry* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  WalEntry* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<WalEntry>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const WalEntry& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const WalEntry& from) {
    WalEntry::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::PROTOBUF_NAMESPACE_ID::Arena* arena, bool is_message_owned);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(WalEntry* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "greptime.v1.WalEntry";
  }
  protected:
  explicit WalEntry(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kMutationsFieldNumber = 1,
  };
  // repeated .greptime.v1.Mutation mutations = 1;
  int mutations_size() const;
  private:
  int _internal_mutations_size() const;
  public:
  void clear_mutations();
  ::greptime::v1::Mutation* mutable_mutations(int index);
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::greptime::v1::Mutation >*
      mutable_mutations();
  private:
  const ::greptime::v1::Mutation& _internal_mutations(int index) const;
  ::greptime::v1::Mutation* _internal_add_mutations();
  public:
  const ::greptime::v1::Mutation& mutations(int index) const;
  ::greptime::v1::Mutation* add_mutations();
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::greptime::v1::Mutation >&
      mutations() const;

  // @@protoc_insertion_point(class_scope:greptime.v1.WalEntry)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::greptime::v1::Mutation > mutations_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_greptime_2fv1_2fwal_2eproto;
};
// -------------------------------------------------------------------

class BytesWrapper final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:greptime.v1.BytesWrapper) */ {
 public:
  inline BytesWrapper() : BytesWrapper(nullptr) {}
  ~BytesWrapper() override;
  explicit PROTOBUF_CONSTEXPR BytesWrapper(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  BytesWrapper(const BytesWrapper& from);
  BytesWrapper(BytesWrapper&& from) noexcept
    : BytesWrapper() {
    *this = ::std::move(from);
  }

  inline BytesWrapper& operator=(const BytesWrapper& from) {
    CopyFrom(from);
    return *this;
  }
  inline BytesWrapper& operator=(BytesWrapper&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const BytesWrapper& default_instance() {
    return *internal_default_instance();
  }
  static inline const BytesWrapper* internal_default_instance() {
    return reinterpret_cast<const BytesWrapper*>(
               &_BytesWrapper_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    2;

  friend void swap(BytesWrapper& a, BytesWrapper& b) {
    a.Swap(&b);
  }
  inline void Swap(BytesWrapper* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(BytesWrapper* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  BytesWrapper* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<BytesWrapper>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const BytesWrapper& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const BytesWrapper& from) {
    BytesWrapper::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::PROTOBUF_NAMESPACE_ID::Arena* arena, bool is_message_owned);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(BytesWrapper* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "greptime.v1.BytesWrapper";
  }
  protected:
  explicit BytesWrapper(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kDataFieldNumber = 1,
  };
  // bytes data = 1;
  void clear_data();
  const std::string& data() const;
  template <typename ArgT0 = const std::string&, typename... ArgT>
  void set_data(ArgT0&& arg0, ArgT... args);
  std::string* mutable_data();
  PROTOBUF_NODISCARD std::string* release_data();
  void set_allocated_data(std::string* data);
  private:
  const std::string& _internal_data() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_data(const std::string& value);
  std::string* _internal_mutable_data();
  public:

  // @@protoc_insertion_point(class_scope:greptime.v1.BytesWrapper)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr data_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_greptime_2fv1_2fwal_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Mutation

// .greptime.v1.OpType op_type = 1;
inline void Mutation::clear_op_type() {
  _impl_.op_type_ = 0;
}
inline ::greptime::v1::OpType Mutation::_internal_op_type() const {
  return static_cast< ::greptime::v1::OpType >(_impl_.op_type_);
}
inline ::greptime::v1::OpType Mutation::op_type() const {
  // @@protoc_insertion_point(field_get:greptime.v1.Mutation.op_type)
  return _internal_op_type();
}
inline void Mutation::_internal_set_op_type(::greptime::v1::OpType value) {
  
  _impl_.op_type_ = value;
}
inline void Mutation::set_op_type(::greptime::v1::OpType value) {
  _internal_set_op_type(value);
  // @@protoc_insertion_point(field_set:greptime.v1.Mutation.op_type)
}

// uint64 sequence = 2;
inline void Mutation::clear_sequence() {
  _impl_.sequence_ = uint64_t{0u};
}
inline uint64_t Mutation::_internal_sequence() const {
  return _impl_.sequence_;
}
inline uint64_t Mutation::sequence() const {
  // @@protoc_insertion_point(field_get:greptime.v1.Mutation.sequence)
  return _internal_sequence();
}
inline void Mutation::_internal_set_sequence(uint64_t value) {
  
  _impl_.sequence_ = value;
}
inline void Mutation::set_sequence(uint64_t value) {
  _internal_set_sequence(value);
  // @@protoc_insertion_point(field_set:greptime.v1.Mutation.sequence)
}

// .greptime.v1.Rows rows = 3;
inline bool Mutation::_internal_has_rows() const {
  return this != internal_default_instance() && _impl_.rows_ != nullptr;
}
inline bool Mutation::has_rows() const {
  return _internal_has_rows();
}
inline const ::greptime::v1::Rows& Mutation::_internal_rows() const {
  const ::greptime::v1::Rows* p = _impl_.rows_;
  return p != nullptr ? *p : reinterpret_cast<const ::greptime::v1::Rows&>(
      ::greptime::v1::_Rows_default_instance_);
}
inline const ::greptime::v1::Rows& Mutation::rows() const {
  // @@protoc_insertion_point(field_get:greptime.v1.Mutation.rows)
  return _internal_rows();
}
inline void Mutation::unsafe_arena_set_allocated_rows(
    ::greptime::v1::Rows* rows) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.rows_);
  }
  _impl_.rows_ = rows;
  if (rows) {
    
  } else {
    
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:greptime.v1.Mutation.rows)
}
inline ::greptime::v1::Rows* Mutation::release_rows() {
  
  ::greptime::v1::Rows* temp = _impl_.rows_;
  _impl_.rows_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old =  reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(temp);
  temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  if (GetArenaForAllocation() == nullptr) { delete old; }
#else  // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return temp;
}
inline ::greptime::v1::Rows* Mutation::unsafe_arena_release_rows() {
  // @@protoc_insertion_point(field_release:greptime.v1.Mutation.rows)
  
  ::greptime::v1::Rows* temp = _impl_.rows_;
  _impl_.rows_ = nullptr;
  return temp;
}
inline ::greptime::v1::Rows* Mutation::_internal_mutable_rows() {
  
  if (_impl_.rows_ == nullptr) {
    auto* p = CreateMaybeMessage<::greptime::v1::Rows>(GetArenaForAllocation());
    _impl_.rows_ = p;
  }
  return _impl_.rows_;
}
inline ::greptime::v1::Rows* Mutation::mutable_rows() {
  ::greptime::v1::Rows* _msg = _internal_mutable_rows();
  // @@protoc_insertion_point(field_mutable:greptime.v1.Mutation.rows)
  return _msg;
}
inline void Mutation::set_allocated_rows(::greptime::v1::Rows* rows) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(_impl_.rows_);
  }
  if (rows) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(rows));
    if (message_arena != submessage_arena) {
      rows = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, rows, submessage_arena);
    }
    
  } else {
    
  }
  _impl_.rows_ = rows;
  // @@protoc_insertion_point(field_set_allocated:greptime.v1.Mutation.rows)
}

// -------------------------------------------------------------------

// WalEntry

// repeated .greptime.v1.Mutation mutations = 1;
inline int WalEntry::_internal_mutations_size() const {
  return _impl_.mutations_.size();
}
inline int WalEntry::mutations_size() const {
  return _internal_mutations_size();
}
inline void WalEntry::clear_mutations() {
  _impl_.mutations_.Clear();
}
inline ::greptime::v1::Mutation* WalEntry::mutable_mutations(int index) {
  // @@protoc_insertion_point(field_mutable:greptime.v1.WalEntry.mutations)
  return _impl_.mutations_.Mutable(index);
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::greptime::v1::Mutation >*
WalEntry::mutable_mutations() {
  // @@protoc_insertion_point(field_mutable_list:greptime.v1.WalEntry.mutations)
  return &_impl_.mutations_;
}
inline const ::greptime::v1::Mutation& WalEntry::_internal_mutations(int index) const {
  return _impl_.mutations_.Get(index);
}
inline const ::greptime::v1::Mutation& WalEntry::mutations(int index) const {
  // @@protoc_insertion_point(field_get:greptime.v1.WalEntry.mutations)
  return _internal_mutations(index);
}
inline ::greptime::v1::Mutation* WalEntry::_internal_add_mutations() {
  return _impl_.mutations_.Add();
}
inline ::greptime::v1::Mutation* WalEntry::add_mutations() {
  ::greptime::v1::Mutation* _add = _internal_add_mutations();
  // @@protoc_insertion_point(field_add:greptime.v1.WalEntry.mutations)
  return _add;
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField< ::greptime::v1::Mutation >&
WalEntry::mutations() const {
  // @@protoc_insertion_point(field_list:greptime.v1.WalEntry.mutations)
  return _impl_.mutations_;
}

// -------------------------------------------------------------------

// BytesWrapper

// bytes data = 1;
inline void BytesWrapper::clear_data() {
  _impl_.data_.ClearToEmpty();
}
inline const std::string& BytesWrapper::data() const {
  // @@protoc_insertion_point(field_get:greptime.v1.BytesWrapper.data)
  return _internal_data();
}
template <typename ArgT0, typename... ArgT>
inline PROTOBUF_ALWAYS_INLINE
void BytesWrapper::set_data(ArgT0&& arg0, ArgT... args) {
 
 _impl_.data_.SetBytes(static_cast<ArgT0 &&>(arg0), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:greptime.v1.BytesWrapper.data)
}
inline std::string* BytesWrapper::mutable_data() {
  std::string* _s = _internal_mutable_data();
  // @@protoc_insertion_point(field_mutable:greptime.v1.BytesWrapper.data)
  return _s;
}
inline const std::string& BytesWrapper::_internal_data() const {
  return _impl_.data_.Get();
}
inline void BytesWrapper::_internal_set_data(const std::string& value) {
  
  _impl_.data_.Set(value, GetArenaForAllocation());
}
inline std::string* BytesWrapper::_internal_mutable_data() {
  
  return _impl_.data_.Mutable(GetArenaForAllocation());
}
inline std::string* BytesWrapper::release_data() {
  // @@protoc_insertion_point(field_release:greptime.v1.BytesWrapper.data)
  return _impl_.data_.Release();
}
inline void BytesWrapper::set_allocated_data(std::string* data) {
  if (data != nullptr) {
    
  } else {
    
  }
  _impl_.data_.SetAllocated(data, GetArenaForAllocation());
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.data_.IsDefault()) {
    _impl_.data_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:greptime.v1.BytesWrapper.data)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------

// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace greptime

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::greptime::v1::OpType> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::greptime::v1::OpType>() {
  return ::greptime::v1::OpType_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_greptime_2fv1_2fwal_2eproto
