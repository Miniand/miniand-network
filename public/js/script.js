(function(global) {
	global.Miniand = {};

	global.Miniand.postUrl = function(url) {
		var form = $('form');
		form.attr('action', url);
		form.attr('method', 'POST');
		console.log(postUrl);
	};
	global.Miniand.setBackgroundToHue = function(hue) {
		$('body').css('background-color', 'hsl(' + hue + ',61%,55%)');
	};
	global.Miniand.parseLinks = function() {
		$('a[data-confirm]').each(function() {
			var el = $(this);
			var confirmText = el.attr('data-confirm');
			el.removeAttr('data-confirm');
			el.click(function(e) {
				if (!confirm(confirmText)) {
					e.preventDefault();
					e.stopImmediatePropagation();
				}
			});
		});
		$('a[data-post]').each(function() {
			var el = $(this);
			var url = el.attr('data-post');
			el.removeAttr('data-post');
			el.click(function(e) {
				e.preventDefault();
				var form = $('form');
				form.attr('method', 'POST');
				form.attr('action', url);
				form.submit();
			});
		});
	}
}(this));

$(document).ready(function() {
	Miniand.parseLinks();
});
