define(['react', 'morearty', 'moment'], function(React, Morearty, moment) {
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
            var time = moment(+moment.utc(binding.get('time'))).format('h:mm:ss')
            return (
                <div className={classes} onClick={this.props.onClick}>
                    <div className="right">{binding.get('ip')}</div>
                    <div className="title">{time}</div>
                    <div className="subtitle">{binding.get('type')}</div>
                    {children}
                </div>
            );
        },
    });
    return PingListItem;
});
