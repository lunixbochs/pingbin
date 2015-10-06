define(['react', 'morearty'], function(React, Morearty) {
    return React.createClass({
        displayName: 'TopBar',
        mixins: [Morearty.Mixin],
        render: function() {
            return (
                <div className="top-bar" ref="topbar">
                    <button>pingb.in</button>
                </div>
            );
        },
    });
});
