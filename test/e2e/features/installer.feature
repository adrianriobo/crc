@installer 
Feature: Installer

    Test CRC installer 

    @darwin
    Scenario: Install CRC 
        Given a environment where CRC is not installed
        When install CRC from installer
        Then CRC is installed

    