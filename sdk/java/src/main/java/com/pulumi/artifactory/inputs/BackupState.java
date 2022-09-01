// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BackupState extends com.pulumi.resources.ResourceArgs {

    public static final BackupState Empty = new BackupState();

    /**
     * If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
     * 
     */
    @Import(name="createArchive")
    private @Nullable Output<Boolean> createArchive;

    /**
     * @return If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> createArchive() {
        return Optional.ofNullable(this.createArchive);
    }

    /**
     * A valid CRON expression that you can use to control backup frequency. Eg: &#34;0 0 12 * * ? &#34;.
     * 
     */
    @Import(name="cronExp")
    private @Nullable Output<String> cronExp;

    /**
     * @return A valid CRON expression that you can use to control backup frequency. Eg: &#34;0 0 12 * * ? &#34;.
     * 
     */
    public Optional<Output<String>> cronExp() {
        return Optional.ofNullable(this.cronExp);
    }

    /**
     * Flag to enable or disable the backup config. Default value is `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Flag to enable or disable the backup config. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * When set, new repositories will not be automatically added to the backup. Default value is `false`.
     * 
     */
    @Import(name="excludeNewRepositories")
    private @Nullable Output<Boolean> excludeNewRepositories;

    /**
     * @return When set, new repositories will not be automatically added to the backup. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> excludeNewRepositories() {
        return Optional.ofNullable(this.excludeNewRepositories);
    }

    /**
     * A list of excluded repositories from the backup. Default is empty list.
     * 
     */
    @Import(name="excludedRepositories")
    private @Nullable Output<List<String>> excludedRepositories;

    /**
     * @return A list of excluded repositories from the backup. Default is empty list.
     * 
     */
    public Optional<Output<List<String>>> excludedRepositories() {
        return Optional.ofNullable(this.excludedRepositories);
    }

    /**
     * When set to true, mission control will not be automatically added to the backup. Default value is &#39;false&#39;.
     * 
     */
    @Import(name="exportMissionControl")
    private @Nullable Output<Boolean> exportMissionControl;

    /**
     * @return When set to true, mission control will not be automatically added to the backup. Default value is &#39;false&#39;.
     * 
     */
    public Optional<Output<Boolean>> exportMissionControl() {
        return Optional.ofNullable(this.exportMissionControl);
    }

    /**
     * The unique ID of the artifactory backup config.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The unique ID of the artifactory backup config.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
     * 
     */
    @Import(name="retentionPeriodHours")
    private @Nullable Output<Integer> retentionPeriodHours;

    /**
     * @return The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
     * 
     */
    public Optional<Output<Integer>> retentionPeriodHours() {
        return Optional.ofNullable(this.retentionPeriodHours);
    }

    /**
     * If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
     * 
     */
    @Import(name="sendMailOnError")
    private @Nullable Output<Boolean> sendMailOnError;

    /**
     * @return If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> sendMailOnError() {
        return Optional.ofNullable(this.sendMailOnError);
    }

    /**
     * If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
     * 
     */
    @Import(name="verifyDiskSpace")
    private @Nullable Output<Boolean> verifyDiskSpace;

    /**
     * @return If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
     * 
     */
    public Optional<Output<Boolean>> verifyDiskSpace() {
        return Optional.ofNullable(this.verifyDiskSpace);
    }

    private BackupState() {}

    private BackupState(BackupState $) {
        this.createArchive = $.createArchive;
        this.cronExp = $.cronExp;
        this.enabled = $.enabled;
        this.excludeNewRepositories = $.excludeNewRepositories;
        this.excludedRepositories = $.excludedRepositories;
        this.exportMissionControl = $.exportMissionControl;
        this.key = $.key;
        this.retentionPeriodHours = $.retentionPeriodHours;
        this.sendMailOnError = $.sendMailOnError;
        this.verifyDiskSpace = $.verifyDiskSpace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BackupState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BackupState $;

        public Builder() {
            $ = new BackupState();
        }

        public Builder(BackupState defaults) {
            $ = new BackupState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createArchive If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder createArchive(@Nullable Output<Boolean> createArchive) {
            $.createArchive = createArchive;
            return this;
        }

        /**
         * @param createArchive If set, backups will be created within a Zip archive (Slow and CPU intensive). Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder createArchive(Boolean createArchive) {
            return createArchive(Output.of(createArchive));
        }

        /**
         * @param cronExp A valid CRON expression that you can use to control backup frequency. Eg: &#34;0 0 12 * * ? &#34;.
         * 
         * @return builder
         * 
         */
        public Builder cronExp(@Nullable Output<String> cronExp) {
            $.cronExp = cronExp;
            return this;
        }

        /**
         * @param cronExp A valid CRON expression that you can use to control backup frequency. Eg: &#34;0 0 12 * * ? &#34;.
         * 
         * @return builder
         * 
         */
        public Builder cronExp(String cronExp) {
            return cronExp(Output.of(cronExp));
        }

        /**
         * @param enabled Flag to enable or disable the backup config. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Flag to enable or disable the backup config. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param excludeNewRepositories When set, new repositories will not be automatically added to the backup. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder excludeNewRepositories(@Nullable Output<Boolean> excludeNewRepositories) {
            $.excludeNewRepositories = excludeNewRepositories;
            return this;
        }

        /**
         * @param excludeNewRepositories When set, new repositories will not be automatically added to the backup. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder excludeNewRepositories(Boolean excludeNewRepositories) {
            return excludeNewRepositories(Output.of(excludeNewRepositories));
        }

        /**
         * @param excludedRepositories A list of excluded repositories from the backup. Default is empty list.
         * 
         * @return builder
         * 
         */
        public Builder excludedRepositories(@Nullable Output<List<String>> excludedRepositories) {
            $.excludedRepositories = excludedRepositories;
            return this;
        }

        /**
         * @param excludedRepositories A list of excluded repositories from the backup. Default is empty list.
         * 
         * @return builder
         * 
         */
        public Builder excludedRepositories(List<String> excludedRepositories) {
            return excludedRepositories(Output.of(excludedRepositories));
        }

        /**
         * @param excludedRepositories A list of excluded repositories from the backup. Default is empty list.
         * 
         * @return builder
         * 
         */
        public Builder excludedRepositories(String... excludedRepositories) {
            return excludedRepositories(List.of(excludedRepositories));
        }

        /**
         * @param exportMissionControl When set to true, mission control will not be automatically added to the backup. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder exportMissionControl(@Nullable Output<Boolean> exportMissionControl) {
            $.exportMissionControl = exportMissionControl;
            return this;
        }

        /**
         * @param exportMissionControl When set to true, mission control will not be automatically added to the backup. Default value is &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder exportMissionControl(Boolean exportMissionControl) {
            return exportMissionControl(Output.of(exportMissionControl));
        }

        /**
         * @param key The unique ID of the artifactory backup config.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The unique ID of the artifactory backup config.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param retentionPeriodHours The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriodHours(@Nullable Output<Integer> retentionPeriodHours) {
            $.retentionPeriodHours = retentionPeriodHours;
            return this;
        }

        /**
         * @param retentionPeriodHours The number of hours to keep a backup before Artifactory will clean it up to free up disk space. Applicable only to non-incremental backups. Default value is 168 hours ie: 7 days.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriodHours(Integer retentionPeriodHours) {
            return retentionPeriodHours(Output.of(retentionPeriodHours));
        }

        /**
         * @param sendMailOnError If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder sendMailOnError(@Nullable Output<Boolean> sendMailOnError) {
            $.sendMailOnError = sendMailOnError;
            return this;
        }

        /**
         * @param sendMailOnError If set, all Artifactory administrators will be notified by email if any problem is encountered during backup. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder sendMailOnError(Boolean sendMailOnError) {
            return sendMailOnError(Output.of(sendMailOnError));
        }

        /**
         * @param verifyDiskSpace If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
         * 
         * @return builder
         * 
         */
        public Builder verifyDiskSpace(@Nullable Output<Boolean> verifyDiskSpace) {
            $.verifyDiskSpace = verifyDiskSpace;
            return this;
        }

        /**
         * @param verifyDiskSpace If set, Artifactory will verify that the backup target location has enough disk space available to hold the backed up data. If there is not enough space available, Artifactory will abort the backup and write a message in the log file. Applicable only to non-incremental backups.
         * 
         * @return builder
         * 
         */
        public Builder verifyDiskSpace(Boolean verifyDiskSpace) {
            return verifyDiskSpace(Output.of(verifyDiskSpace));
        }

        public BackupState build() {
            return $;
        }
    }

}
