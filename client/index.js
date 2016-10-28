var tipz = angular.module('Tipz', []);

tipz.controller('index', ['$scope', '$http', function ($scope, $parent, $http) { 
    $scope.page = "'home.html'";
}]);

tipz.factory('UserService', function() {
    var setUser = function (userName, userID) {
        this.userName = userName;
        this.userID = userID;
    }
    
    return {
        userName : "",
        userID: "",
        setUser: setUser
    };
});


// All data used for navigation
tipz.factory('NavigationService', function() {
    var setProjectName = function (name) {
        this.projectName = name;
    }
    
    var setPage = function (page) {
        this.page = page;
    }
    
    return {
        projectName: "",
        page : "",
        setProjectName: setProjectName,
        setPage: setPage
    };
});
