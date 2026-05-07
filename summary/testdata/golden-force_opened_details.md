
## Note<br/><sub><sup>ℹ️ All remaining changes, mostly for your information</sub></sup>


### Production usage <sup>🏭</sup>

<details>
  <summary>🟰<sup>1</sup></summary>
  <table>
    <tr><td align="left"><sup>🗄️</sup><a href="http://www.squizlabs.com/php-codesniffer">note-prod_usage-requirement/SAME</a></td><td align="right">3.1.1</td></tr>
  </table>

</details>

### Dev-only usage <sup>🧪</sup>

<details>
  <summary>🟰<sup>4</sup></summary>
  <table>
    <tr><td align="left"><sup>🧰</sup><a href="http://www.squizlabs.com/php-codesniffer">note-dev_only_usage-requirement/SAME</a></td><td align="right">3.1.1</td></tr>
    <tr><td align="left"><sup>🧰</sup><a href="http://www.squizlabs.com/php-codesniffer">note-dev_only_usage-requirement/SAME-2</a></td><td align="right">3.1.1</td></tr>
    <tr><td align="left"><sup>🧰</sup><a href="http://www.squizlabs.com/php-codesniffer">note-dev_only_usage-requirement/SAME-3</a></td><td align="right">3.1.1</td></tr>
    <tr><td align="left"><sup>🧰</sup><a href="http://www.squizlabs.com/php-codesniffer">note-dev_only_usage-requirement/SAME-4</a></td><td align="right">3.1.1</td></tr>
  </table>

</details>
<hr/>


<details>
  <summary>Captions</summary>

  #### Versions

  <table>
    <tr><td align="center">_VERSION_❗</td><td align="left">Version is not semver compliant.<br/>Usually a commit ref or branch.</td></tr>
  </table>


  #### Operations

  <table>
    <tr><td align="center">❓</td><td align="left">Unknown update</td></tr>
    <tr><td align="center">❌</td><td align="left">Removed package</td></tr>
    <tr><td align="center">➕️</td><td align="left">Added package</td></tr>
    <tr><td align="center">🟰</td><td align="left">No change</td></tr>
    <tr><td align="center"><sub><sup>🔺.🔹.🔹</sup></sub></td><td align="left">Major upgrade</td></tr>
    <tr><td align="center"><sub><sup>🔻.🔹.🔹</sup></sub></td><td align="left">Major downgrade</td></tr>
    <tr><td align="center"><sub><sup>🔹.🔺.🔹</sup></sub></td><td align="left">Minor upgrade</td></tr>
    <tr><td align="center"><sub><sup>🔹.🔻.🔹</sup></sub></td><td align="left">Minor downgrade</td></tr>
    <tr><td align="center"><sub><sup>🔹.🔹.🔺</sup></sub></td><td align="left">Patch upgrade</td></tr>
    <tr><td align="center"><sub><sup>🔹.🔹.🔻</sup></sub></td><td align="left">Patch downgrade</td></tr>
    <tr><td align="center"><sub><sup>🔹.🔹.🔹❓</sup></sub></td><td align="left">Extra updated, considered as Unknown update</td></tr>
    <tr><td align="center">❔</td><td align="left">Unmanaged operation</td></tr>
  </table>


  #### Package types

  <table>
    <tr><td align="center">🗄</td><td align="left">Package is explicitly required for production usage</td></tr>
    <tr><td align="center">🧰</td><td align="left">Package is explicitly required for dev-only usage</td></tr>
    <tr><td align="center">🔗️</td><td align="left">Transitive dependency package</td></tr>
    <tr><td align="center">💀</td><td align="left">Package is declared abandoned.<br/>You should replace it.</td></tr>
  </table>


  #### Production vs Dev-only usage

  <table>
    <tr><td align="center">🏭</td><td align="left">Package is available in <b>production</b></td></tr>
    <tr><td align="center">🧪</td><td align="left">Package is only available for <b>dev-only</b>, it won't exist in production</td></tr>
  </table>

  There is a difference between **Usage** and **Requirement**.

  - A **Requirement** can be defined as dev-only or not.
    
    Depending on the manager, there might be dedicated property for each environment in the requirement file.
  - **Usage** however is whether the package is available on production or only for dev-only.
    
    Usually, it's the package lock which provides this information.

  You may require a package for dev-only, but this package may actually be a dependency of one of your requirement for production. In that case, the package you required for dev-only will be displayed in "Production usage" section, because it is actually available in production.
</details>
