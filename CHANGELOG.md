## Change Log

### v1.9.5 (2025/5/19)
- Added the incident/list/active method

### v1.9.4 (2023/11/22)
- Added the PostMortemUrl to the incident response (@gsalingu)

### v1.9.3 (2021/9/22)
- Added the silent flag to the subscriber/add method (@jc-dgit)
- Changed the subscriber/add granular field to be optional (@jc-dgit)

### v1.9.2 (2021/9/6)
- Added LICENSE

### v1.9.1 (2021/5/17)
- Added Microsoft Teams

### v1.9.0 (2020/11/16)
- Fix unmarshal of MaintenanceSingleResponse (@danielb42)

### v1.8.0 (2020/4/13)
- Set module name to be the same as repo name (@drornir)
- Fix StatusPartialServiceDisruption value (@drornir)


### v1.7.0 (2020/2/10)
- Fix date format for Maintenance Schedule test
- Adding message_subject to all incident and maintenance methods

### v1.6.0 (2019/11/8)
- enabled go mod + linting
- removed go version from module
- added State and Status types w/enums
- Linting, formatting and modules (@justcompile)

### v1.5.0 (2019/7/12)
- Fixed typo

### v1.4.0 (2018/11/28)
- Changed variables to proper type (int->str)

### v1.3.0 (2018/4/19)
- Fixed invalid response for Incident and Maintenance single calls
- Added missing structure elements for StatusSummary

### v1.2.0 (2018/2/9)
- Support retrieving single incident/maintenance events. New incident/maintenance methods to fetch list of IDs
- Change /component/status/update to use a single component

### v1.1.0 (2017/12/20)
- Updated for new maintenance/schedule infrastructure_affected
- Updated for new incident/create infrastructure affected

### v1.0.0 (2016/1/6)
- Initial release