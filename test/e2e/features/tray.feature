@tray 
Feature: Tray

    Test the cluster management through the tray

    @darwin
    Scenario: Install tray  
        When install CRC tray
        Then tray should be installed
        And tray icon should be accessible

    @darwin
    Scenario: Start Cluster
        Given fresh tray installation   
        When start the cluster from the tray
        And set the pull secret file 
        Then cluster should be started
        And tray should show cluster as running
        And user should get notified with cluster state as running
        
    # @darwin @evaluate
    # Scenario: Start Cluster with wrong pull secret
    #     Given fresh crc tray installation   
    #     When start the cluster from the tray
    #     And set a wrong pull secret file 
    #     Then should get a pop up warning about the pull secret

    @darwin
    Scenario Outline: Connect the cluster
        Given a running cluster   
        When using copied oc login command for <ocp-user>  
        Then user is connected to the cluster as <ocp-user> 
        #And user should get notified with command copied

    Examples:
            | ocp-user   |
            | kubeadmin |
            | developer |

    @darwin 
    Scenario: Stop Cluster
        Given a running cluster   
        When stop the cluster from the tray 
        Then cluster should be stopped
        And tray should show cluster as stopped
        And user should get notified with cluster state as stopped

    @darwin 
    Scenario: Restart Cluster
        Given a stopped cluster   
        When start the cluster from the tray 
        Then cluster should be started
        And tray should show cluster as running
        And user should get notified with cluster state as running

    # @darwin 
    # Scenario: Delete Cluster
    #     Given a running or stopped cluster   
    #     When delete the cluster from the tray 
    #     Then cluster should be deleted
    #     And tray should show cluster as deleted
    #     And user should get notified with cluster state as deleted

    # @darwin 
    # Scenario: Remove tray
    #     Given a environment with tray installed   
    #     When quit tray 
    #     Then tray should be uninstalled

    