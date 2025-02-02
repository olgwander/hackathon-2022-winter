## Basic info

Project name：cicada

Project approval date：2022-11

## Project Intro

### Intoduction
Cicada is a decentralized encyclopedia and content distribution system. It aims to address the following pain points:

- The conflict between platform owners and core value contributors in traditional platforms over the distribution of benefits.
- The governance conflict between platform owners and core value contributors in traditional platforms.
- The problems of scattered, inaccurate, untimely, and inconvenient information access to blockchain projects.

The ultimate goals of Cicada are:
- To enable users to access valuable, systematic content in a timely manner.
- To give participants the right to share benefits, govern, and oversee.
- To fully community-based, reject officialization of project information, and bind more stakeholders to generate strong consensus.

### The main highlights of the project are:

 * Using the economic system as the foundation and the blockchain project as the entry point, the project aims to develop into a comprehensive encyclopedia of knowledge.
 * Using soul tokens to establish user growth and credit systems, and using economic incentives to solve the problem of data collection and user content contributions.
 * Using on-chain governance and rewards to solve the problem of data quality.

### Technical Architecture

* Input and Display Layer

  * react: Used for user interaction and page data display
  * polkadot.jssdk: Used for data interaction with on-chain business logic
  * ipfs.jssdk: Used for file interaction with IPFS storage
* Blockchain Layer

  * Substrate: Used for the implementation of the chain, with business logic including data binding, mortgage token, subscription data, donation, data input, token reward, soul binding, data substitution, token penalty, data complaint, and data appeal
  * Substrate.gov2: Used for the implementation of governance, with business logic including proposal initiation, governance voting, and proposal execution
* Storage Layer

  * Chain: Used for data (not files) storage, processing data submitted by the input and display layer
  * IPFS: Used for file storage, processing files submitted by the input and display layer
* Data Middle Layer

  * Subquery: Used for caching of on-chain data, pulling data from the chain and caching it in the database, and providing data interface services
  * Java: Used for subscription push, getting data from subquery and processing the data format, and regularly pushing to subscribed users
* Brief Description of Business Process (Data Input)
  * Front-end users input data and submit it to the chain, the on-chain logic processes the result and returns it to the user and triggers a chain storage event, the front-end user obtains the on-chain result, the middle layer listens to the chain event, pulls and caches the data from the chain, and the front-end obtains the cached data from the middle layer for display. Here, the on-chain logic processing includes soul binding and token rewards, and if it is data replacement, complaint, and appeal, it also involves governance functions.

- Project logo

![project-logo](https://user-images.githubusercontent.com/40650475/206896925-6228d89f-31f2-4226-aa0b-82a29f71aa83.png)


## Planned tasks to be completed during the hackathon


**Blockchain End**

* `pallet-cicada`
  * Cicada creation and data structure definition (`fn create_category()` `fn create_label()` `fn create_subject()` `fn create_dimension()` `fn create_content()`)
  * Cicada change functions (`fn modify_category()` `fn modify_label()` `fn modify_subject()``fn modify_dimension()` `fn modify_content()`)
  * Cicada transfer function (`fn transfer()`)
  * Cicada destruction function (`fn burn()`)

**Web End**

 * Homepage
 * User page
 * Query and display page
 * Data input page
 * Data governance page

**Middle Layer**

* Subquery End
  * Listen to chain storage events and pull and cache (`function handleCategoryCreatedEvent()` `function handleLabelCreatedEvent()` `function handleSubjectCreatedEvent()` `function handleDimensionCreatedEvent()` `function handleContentCreatedEvent()`)
  * Provide data interface services for clients and Java end
* Java End
  * Get subscribed users (`function get_subscribe()`)


## Team Member Information

| Name | Role | Role description | GitHub Account 
|------|------|--------|--------|
| chenhua | Leader | architecture design |makefriendwithtime
| rxd | Developer | front-end |raoxiandong
| D | Developer | front-end |dengbingbing-tech
| kongbai | Developer | back-end |lxinyuhn
|zhangdong | UI/UX | UI/UX |
|Bella | PM | materials preparation,etc. |zengbing15
