pragma solidity >= 0.5.0;

contract EventNotification {

    event Notify(string msg);

    function broadCast(string calldata m) external {
        emit Notify(m);
    }
}