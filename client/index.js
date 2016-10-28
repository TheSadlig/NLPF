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
    var setProjectID = function (name) {
        this.projectID = name;
    }
    
    var setPage = function (page) {
        this.page = page;
    }
    
    var setReward = function (reward) {
    	this.reward = reward;
    }
    
    var setProject = function (project) {
    	this.project = project;
    }
    var setUser = function (User) {
    	this.User = User;
    }
    
    return {
        projectID: "",
        project: {},
        page : "",
        reward : {},
        User: {},
        setProjectID: setProjectID,
        setPage: setPage,
        setReward: setReward,
        setProject: setProject,
        setUser: setUser
    };
});
