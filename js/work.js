$("#next").on("click",function(){
  var index = $(".active-image").index();
  $(".active-image").fadeOut(500, function(){
    $(".active-image").removeClass("active-image");
    var newActive = $($(".work-images")[index+1]);
    newActive.fadeIn(500);
    newActive.addClass("active-image");
    $("#indicator").text(index+2);
  });
  if ($(".work-images").length >= 2){
    if (index == $(".work-images").length - 2){
      $("#next").prop("disabled", true);
    }
    if (index == 0){
      $("#prev").prop("disabled", false)
    }
  }
});

$("#prev").on("click",function(){
  var index = $(".active-image").index();
  console.log(index);
  var iFrame = $('.active-image');
  var iFrameSRC = iFrame.attr('src');
  iFrame.attr('src','');
  iFrame.attr('src', iFrameSRC);
  $(".active-image").fadeOut(500, function(){
    $(".active-image").removeClass("active-image");
    var newActive = $($(".work-images")[index-1]);
    newActive.fadeIn(500);
    newActive.addClass("active-image");
    $("#indicator").text(index);
  });
  if ($(".work-images").length >= 2){
    if (index == $(".work-images").length - 1){
      $("#next").prop("disabled", false);
    }
    if (index == 1){
      $("#prev").prop("disabled", true);
    }
  }
});
