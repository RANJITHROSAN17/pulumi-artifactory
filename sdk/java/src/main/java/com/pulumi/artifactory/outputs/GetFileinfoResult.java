// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetFileinfoResult {
    /**
     * @return The time &amp; date when the file was created.
     * 
     */
    private final String created;
    /**
     * @return The user who created the file.
     * 
     */
    private final String createdBy;
    /**
     * @return The URI that can be used to download the file.
     * 
     */
    private final String downloadUri;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return The time &amp; date when the file was last modified.
     * 
     */
    private final String lastModified;
    /**
     * @return The time &amp; date when the file was last updated.
     * 
     */
    private final String lastUpdated;
    /**
     * @return MD5 checksum of the file.
     * 
     */
    private final String md5;
    /**
     * @return The mimetype of the file.
     * 
     */
    private final String mimetype;
    /**
     * @return The user who last modified the file.
     * 
     */
    private final String modifiedBy;
    private final String path;
    private final String repository;
    /**
     * @return SHA1 checksum of the file.
     * 
     */
    private final String sha1;
    /**
     * @return SHA256 checksum of the file.
     * 
     */
    private final String sha256;
    /**
     * @return The size of the file.
     * 
     */
    private final Integer size;

    @CustomType.Constructor
    private GetFileinfoResult(
        @CustomType.Parameter("created") String created,
        @CustomType.Parameter("createdBy") String createdBy,
        @CustomType.Parameter("downloadUri") String downloadUri,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("lastModified") String lastModified,
        @CustomType.Parameter("lastUpdated") String lastUpdated,
        @CustomType.Parameter("md5") String md5,
        @CustomType.Parameter("mimetype") String mimetype,
        @CustomType.Parameter("modifiedBy") String modifiedBy,
        @CustomType.Parameter("path") String path,
        @CustomType.Parameter("repository") String repository,
        @CustomType.Parameter("sha1") String sha1,
        @CustomType.Parameter("sha256") String sha256,
        @CustomType.Parameter("size") Integer size) {
        this.created = created;
        this.createdBy = createdBy;
        this.downloadUri = downloadUri;
        this.id = id;
        this.lastModified = lastModified;
        this.lastUpdated = lastUpdated;
        this.md5 = md5;
        this.mimetype = mimetype;
        this.modifiedBy = modifiedBy;
        this.path = path;
        this.repository = repository;
        this.sha1 = sha1;
        this.sha256 = sha256;
        this.size = size;
    }

    /**
     * @return The time &amp; date when the file was created.
     * 
     */
    public String created() {
        return this.created;
    }
    /**
     * @return The user who created the file.
     * 
     */
    public String createdBy() {
        return this.createdBy;
    }
    /**
     * @return The URI that can be used to download the file.
     * 
     */
    public String downloadUri() {
        return this.downloadUri;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The time &amp; date when the file was last modified.
     * 
     */
    public String lastModified() {
        return this.lastModified;
    }
    /**
     * @return The time &amp; date when the file was last updated.
     * 
     */
    public String lastUpdated() {
        return this.lastUpdated;
    }
    /**
     * @return MD5 checksum of the file.
     * 
     */
    public String md5() {
        return this.md5;
    }
    /**
     * @return The mimetype of the file.
     * 
     */
    public String mimetype() {
        return this.mimetype;
    }
    /**
     * @return The user who last modified the file.
     * 
     */
    public String modifiedBy() {
        return this.modifiedBy;
    }
    public String path() {
        return this.path;
    }
    public String repository() {
        return this.repository;
    }
    /**
     * @return SHA1 checksum of the file.
     * 
     */
    public String sha1() {
        return this.sha1;
    }
    /**
     * @return SHA256 checksum of the file.
     * 
     */
    public String sha256() {
        return this.sha256;
    }
    /**
     * @return The size of the file.
     * 
     */
    public Integer size() {
        return this.size;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFileinfoResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String created;
        private String createdBy;
        private String downloadUri;
        private String id;
        private String lastModified;
        private String lastUpdated;
        private String md5;
        private String mimetype;
        private String modifiedBy;
        private String path;
        private String repository;
        private String sha1;
        private String sha256;
        private Integer size;

        public Builder() {
    	      // Empty
        }

        public Builder(GetFileinfoResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.created = defaults.created;
    	      this.createdBy = defaults.createdBy;
    	      this.downloadUri = defaults.downloadUri;
    	      this.id = defaults.id;
    	      this.lastModified = defaults.lastModified;
    	      this.lastUpdated = defaults.lastUpdated;
    	      this.md5 = defaults.md5;
    	      this.mimetype = defaults.mimetype;
    	      this.modifiedBy = defaults.modifiedBy;
    	      this.path = defaults.path;
    	      this.repository = defaults.repository;
    	      this.sha1 = defaults.sha1;
    	      this.sha256 = defaults.sha256;
    	      this.size = defaults.size;
        }

        public Builder created(String created) {
            this.created = Objects.requireNonNull(created);
            return this;
        }
        public Builder createdBy(String createdBy) {
            this.createdBy = Objects.requireNonNull(createdBy);
            return this;
        }
        public Builder downloadUri(String downloadUri) {
            this.downloadUri = Objects.requireNonNull(downloadUri);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder lastModified(String lastModified) {
            this.lastModified = Objects.requireNonNull(lastModified);
            return this;
        }
        public Builder lastUpdated(String lastUpdated) {
            this.lastUpdated = Objects.requireNonNull(lastUpdated);
            return this;
        }
        public Builder md5(String md5) {
            this.md5 = Objects.requireNonNull(md5);
            return this;
        }
        public Builder mimetype(String mimetype) {
            this.mimetype = Objects.requireNonNull(mimetype);
            return this;
        }
        public Builder modifiedBy(String modifiedBy) {
            this.modifiedBy = Objects.requireNonNull(modifiedBy);
            return this;
        }
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public Builder repository(String repository) {
            this.repository = Objects.requireNonNull(repository);
            return this;
        }
        public Builder sha1(String sha1) {
            this.sha1 = Objects.requireNonNull(sha1);
            return this;
        }
        public Builder sha256(String sha256) {
            this.sha256 = Objects.requireNonNull(sha256);
            return this;
        }
        public Builder size(Integer size) {
            this.size = Objects.requireNonNull(size);
            return this;
        }        public GetFileinfoResult build() {
            return new GetFileinfoResult(created, createdBy, downloadUri, id, lastModified, lastUpdated, md5, mimetype, modifiedBy, path, repository, sha1, sha256, size);
        }
    }
}
