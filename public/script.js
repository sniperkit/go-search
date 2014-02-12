$(document).ready(function()
{
  $('#searchform').submit(function() {
    return false;
  });
  $("#searchterm").keyup(function(e){
    var q = $("#searchterm").val();
    $.getJSON("search.json",{term: q}, function(data) {
      $("#results").empty();
      $("#results").append('<li class="list-group-item list-group-item-info">' + "Results in <strong>" + data.timing + "</strong></li>");
      data.values.forEach(function(elem,index) {
        simpletext = new RegExp("(" + q + ")","gi");
        hielem = elem.replace(simpletext, "<strong>$1</strong>");
        $("#results").append('<li class="list-group-item">' + hielem + "</li>");
      });
    });
  });
});
