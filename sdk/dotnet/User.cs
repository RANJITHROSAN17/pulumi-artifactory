// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    /// <summary>
    /// Provides an Artifactory user resource. This can be used to create and manage Artifactory users.
    /// 
    /// When the optional attribute `password` is omitted, a random password is generated according to current Artifactory password policy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Artifactory = Pulumi.Artifactory;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Create a new Artifactory user called terraform
    ///     var test_user = new Artifactory.User("test-user", new()
    ///     {
    ///         Email = "test-user@artifactory-terraform.com",
    ///         Groups = new[]
    ///         {
    ///             "logged-in-users",
    ///             "readers",
    ///         },
    ///         Password = "my super secret password",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Users can be imported using their name, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import artifactory:index/user:User test-user myusername
    /// ```
    /// </summary>
    [ArtifactoryResourceType("artifactory:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        /// </summary>
        [Output("admin")]
        public Output<bool?> Admin { get; private set; } = null!;

        /// <summary>
        /// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        /// </summary>
        [Output("disableUiAccess")]
        public Output<bool?> DisableUiAccess { get; private set; } = null!;

        /// <summary>
        /// Email for user.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// List of groups this user is a part of.
        /// - Note: If "groups" attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Output("internalPasswordDisabled")]
        public Output<bool?> InternalPasswordDisabled { get; private set; } = null!;

        /// <summary>
        /// Username for user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password for the user. When omitted, a random password is generated using the following password policy: 10 characters with 1 digit, 1 symbol, with upper and lower case letters.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        /// </summary>
        [Output("profileUpdatable")]
        public Output<bool?> ProfileUpdatable { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("artifactory:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("artifactory:index/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        /// </summary>
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        /// <summary>
        /// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        /// </summary>
        [Input("disableUiAccess")]
        public Input<bool>? DisableUiAccess { get; set; }

        /// <summary>
        /// Email for user.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of groups this user is a part of.
        /// - Note: If "groups" attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Input("internalPasswordDisabled")]
        public Input<bool>? InternalPasswordDisabled { get; set; }

        /// <summary>
        /// Username for user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for the user. When omitted, a random password is generated using the following password policy: 10 characters with 1 digit, 1 symbol, with upper and lower case letters.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        /// </summary>
        [Input("profileUpdatable")]
        public Input<bool>? ProfileUpdatable { get; set; }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
        /// </summary>
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        /// <summary>
        /// When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
        /// </summary>
        [Input("disableUiAccess")]
        public Input<bool>? DisableUiAccess { get; set; }

        /// <summary>
        /// Email for user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// List of groups this user is a part of.
        /// - Note: If "groups" attribute is not specified then user's group membership set to empty. User will not be part of default "readers" group automatically.
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        /// <summary>
        /// When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
        /// </summary>
        [Input("internalPasswordDisabled")]
        public Input<bool>? InternalPasswordDisabled { get; set; }

        /// <summary>
        /// Username for user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for the user. When omitted, a random password is generated using the following password policy: 10 characters with 1 digit, 1 symbol, with upper and lower case letters.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
        /// </summary>
        [Input("profileUpdatable")]
        public Input<bool>? ProfileUpdatable { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
