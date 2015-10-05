define(['react', 'morearty'], function(React, Morearty) {
    var PingListItem;
    PingListItem = React.createClass({
        displayName: 'PingListItem',
        mixins: [Morearty.Mixin],
        shouldComponentUpdateOverride: function (orig, nextProps, nextState) {
            return this.props.selected !== nextProps.selected || orig();
        },
        render: function() {
            var binding = this.getDefaultBinding();
            var cx = React.addons.classSet;
            var classes = cx({
                'list-item': true,
                'list-item-selected': this.props.selected,
                'intercept-item': binding.get('intercept'),
            });
            var status = binding.get('status');
            if (status) {
                classes += ' status-' + status;
            }
            var children;
            if (binding.get('children')) {
                children = [
                    <PingListItem binding={binding.sub('children').sub(0)} />
                ];
            }
            var title = binding.get('method') + ' ' + binding.get('uri');
            return (
                <div className={classes} onClick={this.props.onClick}>
                    <div className="uri">{title}</div>
                    <div className="host">{binding.get('host')}</div>
                    {children}
                </div>
            );
        },
    });
    return PingListItem;
});
