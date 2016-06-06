# RELEASE PLAN

#### Broad aims:

 - build new boxes to run H2 using Puppet (confirm changes with Boyan)
 - upgrade all services in production to latest RC of platform layer
 - setup monitoring
 - release services to run "tech-free city"
 - configure ATL tech-free city

#### Prerequisites

 - merge RC (platform layer, service layer, hailo lib) into master
 - remove Godeps and rebuild everything

#### Who's involved

 - Dave G (coordinating efforts)
 - Dom W, Senior Engineer (key responsibility RMQ)
 - Boyan D (key responsibility puppet/infrastructure)
 - Jono M (key responsibility testing it works)

#### Initial state

```
#LIVE> versions
Running versions 
  com.hailocab.api.v1.experiment            20131007112316 master 
  com.hailocab.api.v1.gamification          20131107113709 master 
  com.hailocab.api.v1.point                 20131002142909 master 
  com.hailocab.api.v2.h2                    20130829145959 master 
  com.hailocab.kernel.binding               20131002143531 master 
  com.hailocab.kernel.discovery             20131002143347 master 
  com.hailocab.kernel.provisioning          20131022182203 master 
  com.hailocab.service.driver-organization  20131202180828 hotfix 
  com.hailocab.service.event-augmenter      20131112114706 master 
  com.hailocab.service.experiment           20131028185128 master 
  com.hailocab.service.gamification         20131107154444 master 
  com.hailocab.service.geocoding            20131031133150 master 
  com.hailocab.service.i18n-store           20130925140117 master 
  com.hailocab.service.idgen                20130828162103 master 
  com.hailocab.service.login                20131028122844 master 
  com.hailocab.service.nsq-to-hailoengine   20131002143925 master 
  com.hailocab.service.points               20131002142924 master 
```

#### Expected final state

@todo list service versions we'll deploy


## Phase 1: kernel (high risk)

Aim: get new `kernel` H2 boxes up and running with the very latest provisioning, binding, discovery, login and config services.

When: Tuesday 3rd December

Steps: 

 - spin up the new "kernel" machines
 - provision config service on new kernel machines
 - install config (H2:BASE plus each region), check it works (query via hshell, query via HTTP)
 - provision new discovery and binding service onto these machines, check they boot by examining logs
 - shutdown discovery in one AZ on old "points" boxes (manually) - expectation that all services in that AZ re-register with new discovery service
 - repeat for other AZs until old discovery service removed
 - remove old binding service from points boxes
 - install H2 config for new login service
 - provision new login service on kernel boxes, check it boots by examining logs

Risks: binding, discovery and provisioning are all kernel services that run the platform - risk of the entire H2 platform grinding to a halt after/during upgrade

Mitigation plan: remove services provisioned on new kernel machines; replace (if removed) exact versions of discovery/binding/login running in production

## Phase 2: monitoring (low risk)

Aim: get the new monitoring system setup and integrated with Zabbix to trigger alerts

When: Wednesday 4th December

 - spin up the new "stats" machines (for non-critical heavy workloads)
 - provision monitoring service on (@todo which?) boxes
 - connect up Zabbix, test it works (introduce failure into non-critical H2 services and ensure alerts triggered)
 - provision trace service on stats boxes and test (configure a percentage all requests to initiate trace)

Risks: could overload NSQ (unlikely) or stats machines with load from processing trace data

Mitigation plan: pause NSQ trace topic

## Phase 3: new "thin" API, incl. new rule definitions in H2 config service (very high risk)

Aim: get the new "thin" API running, with the new ruleset that will allow H2 to launch tech-free cities and the new /rpc method

When: Thursday 5th December

 - install config/ruleset
 - manually install updated deb package on a single H2 API box in each region that we have removed from the LB pool
 - test it works (test customer API calls, driver API calls in each city @todo outline in greater detail)
 - put back into LB, monitor error rates and overall system health for 30 minutes
 - repeat process with the next box until all are upgraded

Risks: the thin API proxies all driver traffic in some cities (@todo enumerate list) - if it breaks, or we get the rules wrong, we could divert real traffic away from the correct H1 destination, causing an outage

Mitigation plan: monitor traffic carefully and be prepared to install the old deb package if any sign of problems

## Phase 4: points (low risk)

Aim: upgrade services running points to the latest version, such that they talk to H2 config service and integrate with monitoring

When: Monday 9th December

 - com.hailocab.api.v1.point
 - com.hailocab.service.nsq-to-hailoengine
 - com.hailocab.service.points

Risks: we could stop points being ingested via H2 if there are any problems, we could overload H2 connections to master if we have a pause and then release of points

Mitigation plan: ensure no city is sending more than 50% of points via H2 prior to release; dial back up to 100% after launch. Pause NSQ if any suspected overloading problems

## Phase 5: experiments and related (medium risk)

Aim: upgrade services running experiments to latest version

When: Monday 9th December

 - com.hailocab.api.v1.experiment
 - com.hailocab.api.v1.gamification
 - com.hailocab.service.event-augmenter
 - com.hailocab.service.experiment
 - com.hailocab.service.gamification
 - com.hailocab.service.i18n-store
 - com.hailocab.service.idgen

Risks: could break experiments

Mitigation plan: rollback service versions


## Phase 6: driver orgs (medium risk)

Aim: upgrade organisation and geocoding service to latest version, with additional assistance from Bizops

When: Monday 9th December

 - com.hailocab.service.driver-organization
 - com.hailocab.service.geocoding

Risks: could break orgs

Mitigation plan: rollback service versions

## Phase 7: driver/job features (low risk)

Aim: deploy all the services for H2 tech-free city at our leisure

When: Tuesday 10th December

@todo list them

Risks: could overload boxes we are installing them onto, causing problems for live services

Mitigation plan: release services one at a time, ensuring they are fully configured and healthy (in monitoring service) before moving onto the next service. Keep a close eye on system metrics between releases - aim to spread out over the whole day

## Phase 8: configure ATL

Aim: install config for ATL (zoning, config, city, profile, profile-flow)

When: Tuesday 10th December

 - @todo list steps

Risks: minimal/none

## Phase 9: proxy customer API traffic (medium risk)

Aim: update DNS record for api-customer.elasticride.com to proxy traffic via H2 "thin" API in order to allow tech-free city launch

When: Wednesday 11th December

 - ensure ruleset in place to correctly route traffic for all H1 cities, test this
 - update DNS for api-customer.elasticride.com and keep an eye on traffic (expected low volume since cities generally use api-customer-dublin.elasticride.com etc)

Risks: could overload API boxes with lots of traffic; could cause outage if rules not configured and traffic doesn't route correctly

Mitigation plan: switch DNS back


