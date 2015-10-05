define(['react', 'morearty', 'immutable',
        'app/Ctx',
        'app/socket',
        'jsx!app/PingList',
        'jsx!app/PingView',
        'jsx!app/TopBar',
    ], function(
        React, Morearty, Immutable,
        Ctx,
        socket,
        PingList,
        PingView,
        TopBar) {

    return React.createClass({
        displayName: 'App',
        mixins: [Morearty.Mixin],
        render: function() {
            var binding = this.getDefaultBinding();
            return (
                <div className="hbox">
                    <TopBar binding={binding} />
                    <div className="row vbox">
                        <PingList binding={binding} />
                        <PingView binding={binding} />
                    </div>
                </div>
            );
        },
    });
});
