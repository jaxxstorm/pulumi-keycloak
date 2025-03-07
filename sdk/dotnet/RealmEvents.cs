// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    /// <summary>
    /// ## # keycloak..RealmEvents
    /// 
    /// Allows for managing Realm Events settings within Keycloak.
    /// 
    /// ### Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             Realm = "test",
    ///         });
    ///         var realmEvents = new Keycloak.RealmEvents("realmEvents", new Keycloak.RealmEventsArgs
    ///         {
    ///             AdminEventsDetailsEnabled = true,
    ///             AdminEventsEnabled = true,
    ///             EnabledEventTypes = 
    ///             {
    ///                 "LOGIN",
    ///                 "LOGOUT",
    ///             },
    ///             EventsEnabled = true,
    ///             EventsExpiration = 3600,
    ///             EventsListeners = 
    ///             {
    ///                 "jboss-logging",
    ///             },
    ///             RealmId = realm.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The name of the realm the event settings apply to.
    /// - `admin_events_enabled` - (Optional) When true, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
    /// - `admin_events_details_enabled` - (Optional) When true, saved admin events will included detailed information for create/update requests. Defaults to `false`.
    /// - `events_enabled` - (Optional) When true, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
    /// - `events_expiration` - (Optional) The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
    /// - `enabled_event_types` - (Optional) The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
    /// - `events_listeners` - (Optional) The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
    /// </summary>
    public partial class RealmEvents : Pulumi.CustomResource
    {
        [Output("adminEventsDetailsEnabled")]
        public Output<bool?> AdminEventsDetailsEnabled { get; private set; } = null!;

        [Output("adminEventsEnabled")]
        public Output<bool?> AdminEventsEnabled { get; private set; } = null!;

        [Output("enabledEventTypes")]
        public Output<ImmutableArray<string>> EnabledEventTypes { get; private set; } = null!;

        [Output("eventsEnabled")]
        public Output<bool?> EventsEnabled { get; private set; } = null!;

        [Output("eventsExpiration")]
        public Output<int?> EventsExpiration { get; private set; } = null!;

        [Output("eventsListeners")]
        public Output<ImmutableArray<string>> EventsListeners { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a RealmEvents resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RealmEvents(string name, RealmEventsArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/realmEvents:RealmEvents", name, args ?? new RealmEventsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RealmEvents(string name, Input<string> id, RealmEventsState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/realmEvents:RealmEvents", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RealmEvents resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RealmEvents Get(string name, Input<string> id, RealmEventsState? state = null, CustomResourceOptions? options = null)
        {
            return new RealmEvents(name, id, state, options);
        }
    }

    public sealed class RealmEventsArgs : Pulumi.ResourceArgs
    {
        [Input("adminEventsDetailsEnabled")]
        public Input<bool>? AdminEventsDetailsEnabled { get; set; }

        [Input("adminEventsEnabled")]
        public Input<bool>? AdminEventsEnabled { get; set; }

        [Input("enabledEventTypes")]
        private InputList<string>? _enabledEventTypes;
        public InputList<string> EnabledEventTypes
        {
            get => _enabledEventTypes ?? (_enabledEventTypes = new InputList<string>());
            set => _enabledEventTypes = value;
        }

        [Input("eventsEnabled")]
        public Input<bool>? EventsEnabled { get; set; }

        [Input("eventsExpiration")]
        public Input<int>? EventsExpiration { get; set; }

        [Input("eventsListeners")]
        private InputList<string>? _eventsListeners;
        public InputList<string> EventsListeners
        {
            get => _eventsListeners ?? (_eventsListeners = new InputList<string>());
            set => _eventsListeners = value;
        }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public RealmEventsArgs()
        {
        }
    }

    public sealed class RealmEventsState : Pulumi.ResourceArgs
    {
        [Input("adminEventsDetailsEnabled")]
        public Input<bool>? AdminEventsDetailsEnabled { get; set; }

        [Input("adminEventsEnabled")]
        public Input<bool>? AdminEventsEnabled { get; set; }

        [Input("enabledEventTypes")]
        private InputList<string>? _enabledEventTypes;
        public InputList<string> EnabledEventTypes
        {
            get => _enabledEventTypes ?? (_enabledEventTypes = new InputList<string>());
            set => _enabledEventTypes = value;
        }

        [Input("eventsEnabled")]
        public Input<bool>? EventsEnabled { get; set; }

        [Input("eventsExpiration")]
        public Input<int>? EventsExpiration { get; set; }

        [Input("eventsListeners")]
        private InputList<string>? _eventsListeners;
        public InputList<string> EventsListeners
        {
            get => _eventsListeners ?? (_eventsListeners = new InputList<string>());
            set => _eventsListeners = value;
        }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public RealmEventsState()
        {
        }
    }
}
