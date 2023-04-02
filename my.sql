CREATE TABLE `Departments` (
  `DepartmentId` int,
  `Name` int,
  `City` int,
  `PhoneNumber` int
);

CREATE TABLE `Positions` (
  `PositionId` int,
  `Title` int,
  `Description` int,
  `Education` int
);

CREATE TABLE `Employees` (
  `EmployeeId` int,
  `Name` varchar(200),
  `Surname` varchar(200),
  `DateOfBirth` data
);

ALTER TABLE `Employees` ADD FOREIGN KEY (`EmployeeId`) REFERENCES `Positions` (`PositionId`);

ALTER TABLE `Positions` ADD FOREIGN KEY (`PositionId`) REFERENCES `Departments` (`Name`);
