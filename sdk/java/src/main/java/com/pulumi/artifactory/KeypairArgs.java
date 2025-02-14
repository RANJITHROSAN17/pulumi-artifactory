// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KeypairArgs extends com.pulumi.resources.ResourceArgs {

    public static final KeypairArgs Empty = new KeypairArgs();

    /**
     * Will be used as a filename when retrieving the public key via REST API.
     * 
     */
    @Import(name="alias", required=true)
    private Output<String> alias;

    /**
     * @return Will be used as a filename when retrieving the public key via REST API.
     * 
     */
    public Output<String> alias() {
        return this.alias;
    }

    /**
     * A unique identifier for the Key Pair record.
     * 
     */
    @Import(name="pairName", required=true)
    private Output<String> pairName;

    /**
     * @return A unique identifier for the Key Pair record.
     * 
     */
    public Output<String> pairName() {
        return this.pairName;
    }

    /**
     * Key Pair type. Supported types - GPG and RSA.
     * 
     */
    @Import(name="pairType", required=true)
    private Output<String> pairType;

    /**
     * @return Key Pair type. Supported types - GPG and RSA.
     * 
     */
    public Output<String> pairType() {
        return this.pairType;
    }

    /**
     * Passphrase will be used to decrypt the private key. Validated server side.
     * 
     */
    @Import(name="passphrase")
    private @Nullable Output<String> passphrase;

    /**
     * @return Passphrase will be used to decrypt the private key. Validated server side.
     * 
     */
    public Optional<Output<String>> passphrase() {
        return Optional.ofNullable(this.passphrase);
    }

    /**
     * Private key. PEM format will be validated.
     * 
     */
    @Import(name="privateKey", required=true)
    private Output<String> privateKey;

    /**
     * @return Private key. PEM format will be validated.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }

    /**
     * Public key. PEM format will be validated.
     * 
     */
    @Import(name="publicKey", required=true)
    private Output<String> publicKey;

    /**
     * @return Public key. PEM format will be validated.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    private KeypairArgs() {}

    private KeypairArgs(KeypairArgs $) {
        this.alias = $.alias;
        this.pairName = $.pairName;
        this.pairType = $.pairType;
        this.passphrase = $.passphrase;
        this.privateKey = $.privateKey;
        this.publicKey = $.publicKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KeypairArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KeypairArgs $;

        public Builder() {
            $ = new KeypairArgs();
        }

        public Builder(KeypairArgs defaults) {
            $ = new KeypairArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param alias Will be used as a filename when retrieving the public key via REST API.
         * 
         * @return builder
         * 
         */
        public Builder alias(Output<String> alias) {
            $.alias = alias;
            return this;
        }

        /**
         * @param alias Will be used as a filename when retrieving the public key via REST API.
         * 
         * @return builder
         * 
         */
        public Builder alias(String alias) {
            return alias(Output.of(alias));
        }

        /**
         * @param pairName A unique identifier for the Key Pair record.
         * 
         * @return builder
         * 
         */
        public Builder pairName(Output<String> pairName) {
            $.pairName = pairName;
            return this;
        }

        /**
         * @param pairName A unique identifier for the Key Pair record.
         * 
         * @return builder
         * 
         */
        public Builder pairName(String pairName) {
            return pairName(Output.of(pairName));
        }

        /**
         * @param pairType Key Pair type. Supported types - GPG and RSA.
         * 
         * @return builder
         * 
         */
        public Builder pairType(Output<String> pairType) {
            $.pairType = pairType;
            return this;
        }

        /**
         * @param pairType Key Pair type. Supported types - GPG and RSA.
         * 
         * @return builder
         * 
         */
        public Builder pairType(String pairType) {
            return pairType(Output.of(pairType));
        }

        /**
         * @param passphrase Passphrase will be used to decrypt the private key. Validated server side.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(@Nullable Output<String> passphrase) {
            $.passphrase = passphrase;
            return this;
        }

        /**
         * @param passphrase Passphrase will be used to decrypt the private key. Validated server side.
         * 
         * @return builder
         * 
         */
        public Builder passphrase(String passphrase) {
            return passphrase(Output.of(passphrase));
        }

        /**
         * @param privateKey Private key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(Output<String> privateKey) {
            $.privateKey = privateKey;
            return this;
        }

        /**
         * @param privateKey Private key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder privateKey(String privateKey) {
            return privateKey(Output.of(privateKey));
        }

        /**
         * @param publicKey Public key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey Public key. PEM format will be validated.
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        public KeypairArgs build() {
            $.alias = Objects.requireNonNull($.alias, "expected parameter 'alias' to be non-null");
            $.pairName = Objects.requireNonNull($.pairName, "expected parameter 'pairName' to be non-null");
            $.pairType = Objects.requireNonNull($.pairType, "expected parameter 'pairType' to be non-null");
            $.privateKey = Objects.requireNonNull($.privateKey, "expected parameter 'privateKey' to be non-null");
            $.publicKey = Objects.requireNonNull($.publicKey, "expected parameter 'publicKey' to be non-null");
            return $;
        }
    }

}
