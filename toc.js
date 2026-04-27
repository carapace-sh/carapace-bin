// Populate the sidebar
//
// This is a script, and not included directly in the page, to control the total size of the book.
// The TOC contains an entry for each page, so if each page includes a copy of the TOC,
// the total size of the page becomes O(n**2).
class MDBookSidebarScrollbox extends HTMLElement {
    constructor() {
        super();
    }
    connectedCallback() {
        this.innerHTML = '<ol class="chapter"><li class="chapter-item "><a href="carapace-bin.html"><strong aria-hidden="true">1.</strong> carapace-bin</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="install.html"><strong aria-hidden="true">1.1.</strong> Install</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="install/selfupdate.html"><strong aria-hidden="true">1.1.1.</strong> Selfupdate</a></li></ol></li><li class="chapter-item "><a href="setup.html"><strong aria-hidden="true">1.2.</strong> Setup</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="setup/environment.html"><strong aria-hidden="true">1.2.1.</strong> Environment</a></li><li class="chapter-item "><a href="setup/userConfigDir.html"><strong aria-hidden="true">1.2.2.</strong> UserConfigDir</a></li></ol></li><li class="chapter-item "><a href="groups.html"><strong aria-hidden="true">1.3.</strong> Groups</a></li><li class="chapter-item "><a href="variants.html"><strong aria-hidden="true">1.4.</strong> Variants</a></li><li class="chapter-item "><a href="choices.html"><strong aria-hidden="true">1.5.</strong> Choices</a></li><li class="chapter-item "><a href="completers.html"><strong aria-hidden="true">1.6.</strong> Completers</a></li><li class="chapter-item "><a href="style.html"><strong aria-hidden="true">1.7.</strong> Style</a></li><li class="chapter-item "><a href="spec.html"><strong aria-hidden="true">1.8.</strong> Spec</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="spec/user.html"><strong aria-hidden="true">1.8.1.</strong> User</a></li><li class="chapter-item "><a href="spec/bridge.html"><strong aria-hidden="true">1.8.2.</strong> Bridge</a></li><li class="chapter-item "><a href="spec/embed.html"><strong aria-hidden="true">1.8.3.</strong> Embed</a></li><li class="chapter-item "><a href="spec/run.html"><strong aria-hidden="true">1.8.4.</strong> Run</a></li><li class="chapter-item "><a href="spec/shim.html"><strong aria-hidden="true">1.8.5.</strong> Shim</a></li><li class="chapter-item "><a href="spec/scrape.html"><strong aria-hidden="true">1.8.6.</strong> Scrape</a></li><li class="chapter-item "><a href="spec/codegen.html"><strong aria-hidden="true">1.8.7.</strong> Codegen</a></li><li class="chapter-item "><a href="spec/examples.html"><strong aria-hidden="true">1.8.8.</strong> Examples</a></li><li class="chapter-item "><a href="spec/macros.html"><strong aria-hidden="true">1.8.9.</strong> Macros</a></li></ol></li><li class="chapter-item "><a href="overlay.html"><strong aria-hidden="true">1.9.</strong> Overlay</a></li><li class="chapter-item "><a href="variable.html"><strong aria-hidden="true">1.10.</strong> Variable</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="variable/conditions.html"><strong aria-hidden="true">1.10.1.</strong> Conditions</a></li></ol></li><li class="chapter-item "><a href="development.html"><strong aria-hidden="true">1.11.</strong> Development</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="development/projectLayout.html"><strong aria-hidden="true">1.11.1.</strong> Project Layout</a></li><li class="chapter-item "><a href="development/build.html"><strong aria-hidden="true">1.11.2.</strong> Build</a></li><li class="chapter-item "><a href="development/docker.html"><strong aria-hidden="true">1.11.3.</strong> Docker</a></li><li class="chapter-item "><a href="development/creatingCompleters.html"><strong aria-hidden="true">1.11.4.</strong> Creating Completers</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="development/creatingCompleters/manually.html"><strong aria-hidden="true">1.11.4.1.</strong> Manually</a></li><li class="chapter-item "><a href="development/creatingCompleters/parsing.html"><strong aria-hidden="true">1.11.4.2.</strong> Parsing</a></li><li class="chapter-item "><a href="development/creatingCompleters/scraping.html"><strong aria-hidden="true">1.11.4.3.</strong> Scraping</a></li><li class="chapter-item "><a href="development/creatingCompleters/examples.html"><strong aria-hidden="true">1.11.4.4.</strong> Examples</a></li></ol></li><li class="chapter-item "><a href="development/updatingCompleters.html"><strong aria-hidden="true">1.11.5.</strong> Updating Completers</a></li><li class="chapter-item "><a href="development/bestPractices.html"><strong aria-hidden="true">1.11.6.</strong> Best Practices</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="development/bestPractices/coupledActions.html"><strong aria-hidden="true">1.11.6.1.</strong> Coupled Actions</a></li></ol></li><li class="chapter-item "><a href="development/tools.html"><strong aria-hidden="true">1.11.7.</strong> Tools</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="development/tools/carapace-fmt.html"><strong aria-hidden="true">1.11.7.1.</strong> carapace-fmt</a></li><li class="chapter-item "><a href="development/tools/carapace-lint.html"><strong aria-hidden="true">1.11.7.2.</strong> carapace-lint</a></li><li class="chapter-item "><a href="development/tools/carapace-parse.html"><strong aria-hidden="true">1.11.7.3.</strong> carapace-parse</a></li><li class="chapter-item "><a href="development/tools/carapace-generate.html"><strong aria-hidden="true">1.11.7.4.</strong> carapace-generate</a></li></ol></li></ol></li><li class="chapter-item "><a href="release_notes.html"><strong aria-hidden="true">1.12.</strong> Release Notes</a><a class="toggle"><div>❱</div></a></li><li><ol class="section"><li class="chapter-item "><a href="release_notes/v1.x.html"><strong aria-hidden="true">1.12.1.</strong> v1.x</a></li><li class="chapter-item "><a href="release_notes/v1.6.html"><strong aria-hidden="true">1.12.2.</strong> v1.6</a></li><li class="chapter-item "><a href="release_notes/v1.5.html"><strong aria-hidden="true">1.12.3.</strong> v1.5</a></li><li class="chapter-item "><a href="release_notes/v1.4.html"><strong aria-hidden="true">1.12.4.</strong> v1.4</a></li><li class="chapter-item "><a href="release_notes/v1.3.html"><strong aria-hidden="true">1.12.5.</strong> v1.3</a></li><li class="chapter-item "><a href="release_notes/v1.2.html"><strong aria-hidden="true">1.12.6.</strong> v1.2</a></li><li class="chapter-item "><a href="release_notes/v1.1.html"><strong aria-hidden="true">1.12.7.</strong> v1.1</a></li><li class="chapter-item "><a href="release_notes/v1.0.html"><strong aria-hidden="true">1.12.8.</strong> v1.0</a></li><li class="chapter-item "><a href="release_notes/v0.30.html"><strong aria-hidden="true">1.12.9.</strong> v0.30</a></li><li class="chapter-item "><a href="release_notes/v0.29.html"><strong aria-hidden="true">1.12.10.</strong> v0.29</a></li><li class="chapter-item "><a href="release_notes/v0.28.html"><strong aria-hidden="true">1.12.11.</strong> v0.28</a></li><li class="chapter-item "><a href="release_notes/v0.27.html"><strong aria-hidden="true">1.12.12.</strong> v0.27</a></li><li class="chapter-item "><a href="release_notes/v0.26.html"><strong aria-hidden="true">1.12.13.</strong> v0.26</a></li><li class="chapter-item "><a href="release_notes/v0.25.html"><strong aria-hidden="true">1.12.14.</strong> v0.25</a></li><li class="chapter-item "><a href="release_notes/v0.24.html"><strong aria-hidden="true">1.12.15.</strong> v0.24</a></li><li class="chapter-item "><a href="release_notes/v0.23.html"><strong aria-hidden="true">1.12.16.</strong> v0.23</a></li><li class="chapter-item "><a href="release_notes/v0.22.html"><strong aria-hidden="true">1.12.17.</strong> v0.22</a></li><li class="chapter-item "><a href="release_notes/v0.21.html"><strong aria-hidden="true">1.12.18.</strong> v0.21</a></li><li class="chapter-item "><a href="release_notes/v0.20.html"><strong aria-hidden="true">1.12.19.</strong> v0.20</a></li><li class="chapter-item "><a href="release_notes/v0.19.html"><strong aria-hidden="true">1.12.20.</strong> v0.19</a></li><li class="chapter-item "><a href="release_notes/v0.18.html"><strong aria-hidden="true">1.12.21.</strong> v0.18</a></li><li class="chapter-item "><a href="release_notes/v0.17.html"><strong aria-hidden="true">1.12.22.</strong> v0.17</a></li><li class="chapter-item "><a href="release_notes/v0.16.html"><strong aria-hidden="true">1.12.23.</strong> v0.16</a></li><li class="chapter-item "><a href="release_notes/v0.15.html"><strong aria-hidden="true">1.12.24.</strong> v0.15</a></li><li class="chapter-item "><a href="release_notes/v0.14.html"><strong aria-hidden="true">1.12.25.</strong> v0.14</a></li><li class="chapter-item "><a href="release_notes/v0.13.html"><strong aria-hidden="true">1.12.26.</strong> v0.13</a></li><li class="chapter-item "><a href="release_notes/v0.12.html"><strong aria-hidden="true">1.12.27.</strong> v0.12</a></li><li class="chapter-item "><a href="release_notes/v0.11.html"><strong aria-hidden="true">1.12.28.</strong> v0.11</a></li></ol></li></ol></li></ol>';
        // Set the current, active page, and reveal it if it's hidden
        let current_page = document.location.href.toString().split("#")[0].split("?")[0];
        if (current_page.endsWith("/")) {
            current_page += "index.html";
        }
        var links = Array.prototype.slice.call(this.querySelectorAll("a"));
        var l = links.length;
        for (var i = 0; i < l; ++i) {
            var link = links[i];
            var href = link.getAttribute("href");
            if (href && !href.startsWith("#") && !/^(?:[a-z+]+:)?\/\//.test(href)) {
                link.href = path_to_root + href;
            }
            // The "index" page is supposed to alias the first chapter in the book.
            if (link.href === current_page || (i === 0 && path_to_root === "" && current_page.endsWith("/index.html"))) {
                link.classList.add("active");
                var parent = link.parentElement;
                if (parent && parent.classList.contains("chapter-item")) {
                    parent.classList.add("expanded");
                }
                while (parent) {
                    if (parent.tagName === "LI" && parent.previousElementSibling) {
                        if (parent.previousElementSibling.classList.contains("chapter-item")) {
                            parent.previousElementSibling.classList.add("expanded");
                        }
                    }
                    parent = parent.parentElement;
                }
            }
        }
        // Track and set sidebar scroll position
        this.addEventListener('click', function(e) {
            if (e.target.tagName === 'A') {
                sessionStorage.setItem('sidebar-scroll', this.scrollTop);
            }
        }, { passive: true });
        var sidebarScrollTop = sessionStorage.getItem('sidebar-scroll');
        sessionStorage.removeItem('sidebar-scroll');
        if (sidebarScrollTop) {
            // preserve sidebar scroll position when navigating via links within sidebar
            this.scrollTop = sidebarScrollTop;
        } else {
            // scroll sidebar to current active section when navigating via "next/previous chapter" buttons
            var activeSection = document.querySelector('#sidebar .active');
            if (activeSection) {
                activeSection.scrollIntoView({ block: 'center' });
            }
        }
        // Toggle buttons
        var sidebarAnchorToggles = document.querySelectorAll('#sidebar a.toggle');
        function toggleSection(ev) {
            ev.currentTarget.parentElement.classList.toggle('expanded');
        }
        Array.from(sidebarAnchorToggles).forEach(function (el) {
            el.addEventListener('click', toggleSection);
        });
    }
}
window.customElements.define("mdbook-sidebar-scrollbox", MDBookSidebarScrollbox);
