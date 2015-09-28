/**
 * Created with IntelliJ IDEA.
 * User: zhuclude
 * Date: 6/20/13
 * Time: 1:46 PM
 * To change this template use File | Settings | File Templates.
 */

/* Services */
angular.module('ni.services', [])
    .factory('NiHttp', ['$http','$q', function ($http,$q) {
        return {
            sendRequest: function(options){
                var deferred = $q.defer(),
                    promise = deferred.promise;

                promise.success = function(fn) {
                    promise.then(function(result) {
                        fn(result);
                    });
                    return promise;
                };

                promise.error = function(fn) {
                    promise.then(null, function(result) {
                        fn(result);
                    });
                    return promise;
                };

                var config = {
                    method:         'GET',
                    url:            '',
                    params:         null,
                    data:           null,
                    headers:        null,
                    successFn:      null,
                    errorFn:        null,
                    waitingElement: null,
                    //data:           '',
                    isNeedShowWaiting: false //true
                };
                var opt = angular.extend(config, options);
                // TODO: show block on waitingElement
                if(opt.isNeedShowWaiting){
                    // VX.UI.showWaiting(null, true);
                }

                $http(
                    {
                        method:     opt.method,
                        url:        opt.url,
                        data:       opt.data,
                        headers:    opt.headers,
                        params:     opt.params
                    }
                ).
                success(function(data, status, headers, config) {
                    // if(opt.isNeedShowWaiting) VX.UI.showWaiting(null, false);
                    if(data.status == 1){
                        deferred.resolve(data);
                    }else{
                        deferred.reject(data);
                    }
                }).
                error(function(data, status, headers, config) {
                    // if(opt.isNeedShowWaiting) VX.UI.showWaiting(null, false);
                    deferred.reject(data);
                });

                return promise;
            },

            post: function(options, srcOptions){
                options = options || {};
                var defaultOpt = {method: 'POST', headers:{'X-Requested-With':'XMLHttpRequest'}}
                var opt = angular.extend(options, defaultOpt);
                if(srcOptions){
                    opt = angular.extend(opt, srcOptions);
                }
                return this.sendRequest(opt);
            },

            get: function(options, srcOptions){
                options = options || {};
                var defaultOpt = {method: 'GET'}
                var opt = angular.extend(options, defaultOpt);
                if(srcOptions){
                    opt = angular.extend(opt, srcOptions);
                }
                return this.sendRequest(opt);
            }
        };
    }])
    .factory('Api', ['NiHttp', function (NiHttp) {
        return {
            getRecord: function(pSet, pKey){
                var opts = {
                    url: '/api/de/record',
                    params: {
                        namespace: 'camel',
                        set: pSet,
                        key: pKey
                    }
                };
                return NiHttp.get(opts)
            }
        }
    }])
;