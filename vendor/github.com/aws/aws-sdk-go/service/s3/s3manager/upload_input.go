// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3manager

import (
	"io"
	"time"
)

// UploadInput provides the input parameters for uploading a stream or buffer
// to an object in an Amazon S3 bucket. This type is similar to the s3
// package's PutObjectInput with the exception that the Body member is an
// io.Reader instead of an io.ReadSeeker.
//
// The ContentMD5 member for pre-computed MD5 checksums will be ignored for
// multipart uploads. Objects that will be uploaded in a single part, the
// ContentMD5 will be used.
//
// The Checksum members for pre-computed checksums will be ignored for
// multipart uploads. Objects that will be uploaded in a single part, will
// include the checksum member in the request.
type UploadInput struct {
	_ struct{} `locationName:"PutObjectRequest" type:"structure" payload:"Body"`

	// The canned ACL to apply to the object. For more information, see Canned ACL
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
	//
	// This action is not supported by Amazon S3 on Outposts.
	ACL *string `location:"header" locationName:"x-amz-acl" type:"string" enum:"ObjectCannedACL"`

	// The readable body payload to send to S3.
	Body io.Reader

	// The bucket name to which the PUT action was initiated.
	//
	// When using this action with an access point, you must direct requests to
	// the access point hostname. The access point hostname takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this action with an access point through the Amazon Web Services
	// SDKs, you provide the access point ARN in place of the bucket name. For more
	// information about access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide.
	//
	// When you use this action with Amazon S3 on Outposts, you must direct requests
	// to the S3 on Outposts hostname. The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When
	// you use this action with S3 on Outposts through the Amazon Web Services SDKs,
	// you provide the Outposts access point ARN in place of the bucket name. For
	// more information about S3 on Outposts ARNs, see What is S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption
	// with server-side encryption using Key Management Service (KMS) keys (SSE-KMS).
	// Setting this header to true causes Amazon S3 to use an S3 Bucket Key for
	// object encryption with SSE-KMS.
	//
	// Specifying this header with a PUT action doesn’t affect bucket-level settings
	// for S3 Bucket Key.
	BucketKeyEnabled *bool `location:"header" locationName:"x-amz-server-side-encryption-bucket-key-enabled" type:"boolean"`

	// Can be used to specify caching behavior along the request/reply chain. For
	// more information, see http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9
	// (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9).
	CacheControl *string `location:"header" locationName:"Cache-Control" type:"string"`

	// Indicates the algorithm used to create the checksum for the object when using
	// the SDK. This header will not provide any additional functionality if not
	// using the SDK. When sending this header, there must be a corresponding x-amz-checksum
	// or x-amz-trailer header sent. Otherwise, Amazon S3 fails the request with
	// the HTTP status code 400 Bad Request. For more information, see Checking
	// object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide.
	//
	// If you provide an individual checksum, Amazon S3 ignores any provided ChecksumAlgorithm
	// parameter.
	//
	// The AWS SDK for Go v1 does not support automatic computing request payload
	// checksum. This feature is available in the AWS SDK for Go v2. If a value
	// is specified for this parameter, the matching algorithm's checksum member
	// must be populated with the algorithm's checksum of the request payload.
	ChecksumAlgorithm *string `location:"header" locationName:"x-amz-sdk-checksum-algorithm" type:"string" enum:"ChecksumAlgorithm"`

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This header specifies
	// the base64-encoded, 32-bit CRC32 checksum of the object. For more information,
	// see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide.
	ChecksumCRC32 *string `location:"header" locationName:"x-amz-checksum-crc32" type:"string"`

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This header specifies
	// the base64-encoded, 32-bit CRC32C checksum of the object. For more information,
	// see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide.
	ChecksumCRC32C *string `location:"header" locationName:"x-amz-checksum-crc32c" type:"string"`

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This header specifies
	// the base64-encoded, 160-bit SHA-1 digest of the object. For more information,
	// see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide.
	ChecksumSHA1 *string `location:"header" locationName:"x-amz-checksum-sha1" type:"string"`

	// This header can be used as a data integrity check to verify that the data
	// received is the same data that was originally sent. This header specifies
	// the base64-encoded, 256-bit SHA-256 digest of the object. For more information,
	// see Checking object integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide.
	ChecksumSHA256 *string `location:"header" locationName:"x-amz-checksum-sha256" type:"string"`

	// Specifies presentational information for the object. For more information,
	// see https://www.rfc-editor.org/rfc/rfc6266#section-4 (https://www.rfc-editor.org/rfc/rfc6266#section-4).
	ContentDisposition *string `location:"header" locationName:"Content-Disposition" type:"string"`

	// Specifies what content encodings have been applied to the object and thus
	// what decoding mechanisms must be applied to obtain the media-type referenced
	// by the Content-Type header field. For more information, see https://www.rfc-editor.org/rfc/rfc9110.html#field.content-encoding
	// (https://www.rfc-editor.org/rfc/rfc9110.html#field.content-encoding).
	ContentEncoding *string `location:"header" locationName:"Content-Encoding" type:"string"`

	// The language the content is in.
	ContentLanguage *string `location:"header" locationName:"Content-Language" type:"string"`

	// The base64-encoded 128-bit MD5 digest of the message (without the headers)
	// according to RFC 1864. This header can be used as a message integrity check
	// to verify that the data is the same data that was originally sent. Although
	// it is optional, we recommend using the Content-MD5 mechanism as an end-to-end
	// integrity check. For more information about REST request authentication,
	// see REST Authentication (https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html).
	//
	// If the ContentMD5 is provided for a multipart upload, it will be ignored.
	// Objects that will be uploaded in a single part, the ContentMD5 will be used.
	ContentMD5 *string `location:"header" locationName:"Content-MD5" type:"string"`

	// A standard MIME type describing the format of the contents. For more information,
	// see https://www.rfc-editor.org/rfc/rfc9110.html#name-content-type (https://www.rfc-editor.org/rfc/rfc9110.html#name-content-type).
	ContentType *string `location:"header" locationName:"Content-Type" type:"string"`

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string `location:"header" locationName:"x-amz-expected-bucket-owner" type:"string"`

	// The date and time at which the object is no longer cacheable. For more information,
	// see https://www.rfc-editor.org/rfc/rfc7234#section-5.3 (https://www.rfc-editor.org/rfc/rfc7234#section-5.3).
	Expires *time.Time `location:"header" locationName:"Expires" type:"timestamp"`

	// Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantFullControl *string `location:"header" locationName:"x-amz-grant-full-control" type:"string"`

	// Allows grantee to read the object data and its metadata.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantRead *string `location:"header" locationName:"x-amz-grant-read" type:"string"`

	// Allows grantee to read the object ACL.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantReadACP *string `location:"header" locationName:"x-amz-grant-read-acp" type:"string"`

	// Allows grantee to write the ACL for the applicable object.
	//
	// This action is not supported by Amazon S3 on Outposts.
	GrantWriteACP *string `location:"header" locationName:"x-amz-grant-write-acp" type:"string"`

	// Object key for which the PUT action was initiated.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string `location:"headers" locationName:"x-amz-meta-" type:"map"`

	// Specifies whether a legal hold will be applied to this object. For more information
	// about S3 Object Lock, see Object Lock (https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html).
	ObjectLockLegalHoldStatus *string `location:"header" locationName:"x-amz-object-lock-legal-hold" type:"string" enum:"ObjectLockLegalHoldStatus"`

	// The Object Lock mode that you want to apply to this object.
	ObjectLockMode *string `location:"header" locationName:"x-amz-object-lock-mode" type:"string" enum:"ObjectLockMode"`

	// The date and time when you want this object's Object Lock to expire. Must
	// be formatted as a timestamp parameter.
	ObjectLockRetainUntilDate *time.Time `location:"header" locationName:"x-amz-object-lock-retain-until-date" type:"timestamp" timestampFormat:"iso8601"`

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from Requester Pays buckets, see Downloading Objects
	// in Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide.
	RequestPayer *string `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"RequestPayer"`

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string `location:"header" locationName:"x-amz-server-side-encryption-customer-algorithm" type:"string"`

	// Specifies the customer-provided encryption key for Amazon S3 to use in encrypting
	// data. This value is used to store the object and then it is discarded; Amazon
	// S3 does not store the encryption key. The key must be appropriate for use
	// with the algorithm specified in the x-amz-server-side-encryption-customer-algorithm
	// header.
	SSECustomerKey *string `marshal-as:"blob" location:"header" locationName:"x-amz-server-side-encryption-customer-key" type:"string" sensitive:"true"`

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string `location:"header" locationName:"x-amz-server-side-encryption-customer-key-MD5" type:"string"`

	// Specifies the Amazon Web Services KMS Encryption Context to use for object
	// encryption. The value of this header is a base64-encoded UTF-8 string holding
	// JSON with the encryption context key-value pairs. This value is stored as
	// object metadata and automatically gets passed on to Amazon Web Services KMS
	// for future GetObject or CopyObject operations on this object.
	SSEKMSEncryptionContext *string `location:"header" locationName:"x-amz-server-side-encryption-context" type:"string" sensitive:"true"`

	// If x-amz-server-side-encryption has a valid value of aws:kms or aws:kms:dsse,
	// this header specifies the ID of the Key Management Service (KMS) symmetric
	// encryption customer managed key that was used for the object. If you specify
	// x-amz-server-side-encryption:aws:kms or x-amz-server-side-encryption:aws:kms:dsse,
	// but do not providex-amz-server-side-encryption-aws-kms-key-id, Amazon S3
	// uses the Amazon Web Services managed key (aws/s3) to protect the data. If
	// the KMS key does not exist in the same account that's issuing the command,
	// you must use the full ARN and not just the ID.
	SSEKMSKeyId *string `location:"header" locationName:"x-amz-server-side-encryption-aws-kms-key-id" type:"string" sensitive:"true"`

	// The server-side encryption algorithm used when storing this object in Amazon
	// S3 (for example, AES256, aws:kms, aws:kms:dsse).
	ServerSideEncryption *string `location:"header" locationName:"x-amz-server-side-encryption" type:"string" enum:"ServerSideEncryption"`

	// By default, Amazon S3 uses the STANDARD Storage Class to store newly created
	// objects. The STANDARD storage class provides high durability and high availability.
	// Depending on performance needs, you can specify a different Storage Class.
	// Amazon S3 on Outposts only uses the OUTPOSTS Storage Class. For more information,
	// see Storage Classes (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
	// in the Amazon S3 User Guide.
	StorageClass *string `location:"header" locationName:"x-amz-storage-class" type:"string" enum:"StorageClass"`

	// The tag-set for the object. The tag-set must be encoded as URL Query parameters.
	// (For example, "Key1=Value1")
	Tagging *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// If the bucket is configured as a website, redirects requests for this object
	// to another object in the same bucket or to an external URL. Amazon S3 stores
	// the value of this header in the object metadata. For information about object
	// metadata, see Object Key and Metadata (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html).
	//
	// In the following example, the request header sets the redirect to an object
	// (anotherPage.html) in the same bucket:
	//
	// x-amz-website-redirect-location: /anotherPage.html
	//
	// In the following example, the request header sets the object redirect to
	// another website:
	//
	// x-amz-website-redirect-location: http://www.example.com/
	//
	// For more information about website hosting in Amazon S3, see Hosting Websites
	// on Amazon S3 (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html)
	// and How to Configure Website Page Redirects (https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
	WebsiteRedirectLocation *string `location:"header" locationName:"x-amz-website-redirect-location" type:"string"`
}
