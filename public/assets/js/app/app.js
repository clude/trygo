'use strict';

if(window.ni){

} else{
    window.ni = {};
}

String.format = function() {
    if (arguments.length == 0)
        return null;

    var str = arguments[0];
    for (var i = 1; i < arguments.length; i++) {
        var re = new RegExp('\\{' + (i - 1) + '\\}', 'gm');
        str = str.replace(re, arguments[i]);
    }
    return str;
};

ni.NiNgModule = angular.module('ni', ['ui.bootstrap', 'ni.services'])
    .run(['$rootScope', function($rootScope) {
        $rootScope.isEmpty = function(v) {
            return _.isEmpty(v);
        }
    }])
;






