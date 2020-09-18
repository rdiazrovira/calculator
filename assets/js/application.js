require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

$(() => {
    $('.theme-radio').on('click', function(){
        $("#set-theme-form").submit();
    });
});
