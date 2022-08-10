$(function(){
	$('.menuItems li > a').click(function(){
		if($(this).find('.toggle-icon').eq(0).hasClass('glyphicon-menu-left')){
			$(this).find('.toggle-icon').eq(0).removeClass('glyphicon-menu-left').addClass('glyphicon-menu-down');
		}else{
			$(this).find('.toggle-icon').eq(0).removeClass('glyphicon-menu-down').addClass('glyphicon-menu-left');
		}
		$(this).next().slideToggle('normal');
	});
});