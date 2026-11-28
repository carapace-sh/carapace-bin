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
        this.innerHTML = '<ol class="chapter"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="carapace-bin.html"><strong aria-hidden="true">1.</strong> carapace-bin</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="install.html"><strong aria-hidden="true">1.1.</strong> Install</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="install/selfupdate.html"><strong aria-hidden="true">1.1.1.</strong> Selfupdate</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="setup.html"><strong aria-hidden="true">1.2.</strong> Setup</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="setup/environment.html"><strong aria-hidden="true">1.2.1.</strong> Environment</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="setup/userConfigDir.html"><strong aria-hidden="true">1.2.2.</strong> UserConfigDir</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="groups.html"><strong aria-hidden="true">1.3.</strong> Groups</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="variants.html"><strong aria-hidden="true">1.4.</strong> Variants</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="choices.html"><strong aria-hidden="true">1.5.</strong> Choices</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="completers.html"><strong aria-hidden="true">1.6.</strong> Completers</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="style.html"><strong aria-hidden="true">1.7.</strong> Style</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec.html"><strong aria-hidden="true">1.8.</strong> Spec</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/user.html"><strong aria-hidden="true">1.8.1.</strong> User</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/bridge.html"><strong aria-hidden="true">1.8.2.</strong> Bridge</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/embed.html"><strong aria-hidden="true">1.8.3.</strong> Embed</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/run.html"><strong aria-hidden="true">1.8.4.</strong> Run</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/shim.html"><strong aria-hidden="true">1.8.5.</strong> Shim</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/scrape.html"><strong aria-hidden="true">1.8.6.</strong> Scrape</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/codegen.html"><strong aria-hidden="true">1.8.7.</strong> Codegen</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/examples.html"><strong aria-hidden="true">1.8.8.</strong> Examples</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="spec/macros.html"><strong aria-hidden="true">1.8.9.</strong> Macros</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="overlay.html"><strong aria-hidden="true">1.9.</strong> Overlay</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="variable.html"><strong aria-hidden="true">1.10.</strong> Variable</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="variable/conditions.html"><strong aria-hidden="true">1.10.1.</strong> Conditions</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware.html"><strong aria-hidden="true">1.11.</strong> Slopware</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/mcp.html"><strong aria-hidden="true">1.11.1.</strong> MCP</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills.html"><strong aria-hidden="true">1.11.2.</strong> Skills</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/brew.html"><strong aria-hidden="true">1.11.2.1.</strong> brew</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/bubbletea.html"><strong aria-hidden="true">1.11.2.2.</strong> bubbletea</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/bun.html"><strong aria-hidden="true">1.11.2.3.</strong> bun</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/but.html"><strong aria-hidden="true">1.11.2.4.</strong> but</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/carapace.html"><strong aria-hidden="true">1.11.2.5.</strong> carapace</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/cargo.html"><strong aria-hidden="true">1.11.2.6.</strong> cargo</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/cobra.html"><strong aria-hidden="true">1.11.2.7.</strong> cobra</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/dnf5.html"><strong aria-hidden="true">1.11.2.8.</strong> dnf5</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/gh.html"><strong aria-hidden="true">1.11.2.9.</strong> gh</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/git.html"><strong aria-hidden="true">1.11.2.10.</strong> git</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/go.html"><strong aria-hidden="true">1.11.2.11.</strong> go</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/httpie.html"><strong aria-hidden="true">1.11.2.12.</strong> httpie</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/iproute2.html"><strong aria-hidden="true">1.11.2.13.</strong> iproute2</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/kubeadm.html"><strong aria-hidden="true">1.11.2.14.</strong> kubeadm</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/kubectl.html"><strong aria-hidden="true">1.11.2.15.</strong> kubectl</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/lore.html"><strong aria-hidden="true">1.11.2.16.</strong> lore</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/pandoc.html"><strong aria-hidden="true">1.11.2.17.</strong> pandoc</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/qemu.html"><strong aria-hidden="true">1.11.2.18.</strong> qemu</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/rails.html"><strong aria-hidden="true">1.11.2.19.</strong> rails</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/ruby.html"><strong aria-hidden="true">1.11.2.20.</strong> ruby</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/rustup.html"><strong aria-hidden="true">1.11.2.21.</strong> rustup</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/skills/winget.html"><strong aria-hidden="true">1.11.2.22.</strong> winget</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="slopware/lsp.html"><strong aria-hidden="true">1.11.3.</strong> LSP</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development.html"><strong aria-hidden="true">1.12.</strong> Development</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/projectLayout.html"><strong aria-hidden="true">1.12.1.</strong> Project Layout</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/build.html"><strong aria-hidden="true">1.12.2.</strong> Build</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/docker.html"><strong aria-hidden="true">1.12.3.</strong> Docker</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/creatingCompleters.html"><strong aria-hidden="true">1.12.4.</strong> Creating Completers</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/creatingCompleters/manually.html"><strong aria-hidden="true">1.12.4.1.</strong> Manually</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/creatingCompleters/parsing.html"><strong aria-hidden="true">1.12.4.2.</strong> Parsing</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/creatingCompleters/scraping.html"><strong aria-hidden="true">1.12.4.3.</strong> Scraping</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/creatingCompleters/examples.html"><strong aria-hidden="true">1.12.4.4.</strong> Examples</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/updatingCompleters.html"><strong aria-hidden="true">1.12.5.</strong> Updating Completers</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/bestPractices.html"><strong aria-hidden="true">1.12.6.</strong> Best Practices</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/bestPractices/coupledActions.html"><strong aria-hidden="true">1.12.6.1.</strong> Coupled Actions</a></span></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/tools.html"><strong aria-hidden="true">1.12.7.</strong> Tools</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/tools/carapace-lint.html"><strong aria-hidden="true">1.12.7.1.</strong> carapace-lint</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/tools/carapace-parse.html"><strong aria-hidden="true">1.12.7.2.</strong> carapace-parse</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="development/tools/carapace-generate.html"><strong aria-hidden="true">1.12.7.3.</strong> carapace-generate</a></span></li></ol></li></ol><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes.html"><strong aria-hidden="true">1.13.</strong> Release Notes</a><a class="chapter-fold-toggle"><div>❱</div></a></span><ol class="section"><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.x.html"><strong aria-hidden="true">1.13.1.</strong> v1.x</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.8.html"><strong aria-hidden="true">1.13.2.</strong> v1.8</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.7.html"><strong aria-hidden="true">1.13.3.</strong> v1.7</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.6.html"><strong aria-hidden="true">1.13.4.</strong> v1.6</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.5.html"><strong aria-hidden="true">1.13.5.</strong> v1.5</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.4.html"><strong aria-hidden="true">1.13.6.</strong> v1.4</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.3.html"><strong aria-hidden="true">1.13.7.</strong> v1.3</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.2.html"><strong aria-hidden="true">1.13.8.</strong> v1.2</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.1.html"><strong aria-hidden="true">1.13.9.</strong> v1.1</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v1.0.html"><strong aria-hidden="true">1.13.10.</strong> v1.0</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.30.html"><strong aria-hidden="true">1.13.11.</strong> v0.30</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.29.html"><strong aria-hidden="true">1.13.12.</strong> v0.29</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.28.html"><strong aria-hidden="true">1.13.13.</strong> v0.28</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.27.html"><strong aria-hidden="true">1.13.14.</strong> v0.27</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.26.html"><strong aria-hidden="true">1.13.15.</strong> v0.26</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.25.html"><strong aria-hidden="true">1.13.16.</strong> v0.25</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.24.html"><strong aria-hidden="true">1.13.17.</strong> v0.24</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.23.html"><strong aria-hidden="true">1.13.18.</strong> v0.23</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.22.html"><strong aria-hidden="true">1.13.19.</strong> v0.22</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.21.html"><strong aria-hidden="true">1.13.20.</strong> v0.21</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.20.html"><strong aria-hidden="true">1.13.21.</strong> v0.20</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.19.html"><strong aria-hidden="true">1.13.22.</strong> v0.19</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.18.html"><strong aria-hidden="true">1.13.23.</strong> v0.18</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.17.html"><strong aria-hidden="true">1.13.24.</strong> v0.17</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.16.html"><strong aria-hidden="true">1.13.25.</strong> v0.16</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.15.html"><strong aria-hidden="true">1.13.26.</strong> v0.15</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.14.html"><strong aria-hidden="true">1.13.27.</strong> v0.14</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.13.html"><strong aria-hidden="true">1.13.28.</strong> v0.13</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.12.html"><strong aria-hidden="true">1.13.29.</strong> v0.12</a></span></li><li class="chapter-item "><span class="chapter-link-wrapper"><a href="release_notes/v0.11.html"><strong aria-hidden="true">1.13.30.</strong> v0.11</a></span></li></ol></li></ol></li></ol>';
        // Set the current, active page, and reveal it if it's hidden
        let current_page = document.location.href.toString().split('#')[0].split('?')[0];
        if (current_page.endsWith('/')) {
            current_page += 'index.html';
        }
        const links = Array.prototype.slice.call(this.querySelectorAll('a'));
        const l = links.length;
        for (let i = 0; i < l; ++i) {
            const link = links[i];
            const href = link.getAttribute('href');
            if (href && !href.startsWith('#') && !/^(?:[a-z+]+:)?\/\//.test(href)) {
                link.href = path_to_root + href;
            }
            // The 'index' page is supposed to alias the first chapter in the book.
            // Check both with and without the '.html' suffix to be robust against pretty URLs
            if (link.href.replace(/\.html$/, '') === current_page.replace(/\.html$/, '')
                || i === 0
                && path_to_root === ''
                && current_page.endsWith('/index.html')) {
                link.classList.add('active');
                let parent = link.parentElement;
                while (parent) {
                    if (parent.tagName === 'LI' && parent.classList.contains('chapter-item')) {
                        parent.classList.add('expanded');
                    }
                    parent = parent.parentElement;
                }
            }
        }
        // Track and set sidebar scroll position
        this.addEventListener('click', e => {
            if (e.target.tagName === 'A') {
                const clientRect = e.target.getBoundingClientRect();
                const sidebarRect = this.getBoundingClientRect();
                sessionStorage.setItem('sidebar-scroll-offset', clientRect.top - sidebarRect.top);
            }
        }, { passive: true });
        const sidebarScrollOffset = sessionStorage.getItem('sidebar-scroll-offset');
        sessionStorage.removeItem('sidebar-scroll-offset');
        if (sidebarScrollOffset !== null) {
            // preserve sidebar scroll position when navigating via links within sidebar
            const activeSection = this.querySelector('.active');
            if (activeSection) {
                const clientRect = activeSection.getBoundingClientRect();
                const sidebarRect = this.getBoundingClientRect();
                const currentOffset = clientRect.top - sidebarRect.top;
                this.scrollTop += currentOffset - parseFloat(sidebarScrollOffset);
            }
        } else {
            // scroll sidebar to current active section when navigating via
            // 'next/previous chapter' buttons
            const activeSection = document.querySelector('#mdbook-sidebar .active');
            if (activeSection) {
                activeSection.scrollIntoView({ block: 'center' });
            }
        }
        // Toggle buttons
        const sidebarAnchorToggles = document.querySelectorAll('.chapter-fold-toggle');
        function toggleSection(ev) {
            ev.currentTarget.parentElement.parentElement.classList.toggle('expanded');
        }
        Array.from(sidebarAnchorToggles).forEach(el => {
            el.addEventListener('click', toggleSection);
        });
    }
}
window.customElements.define('mdbook-sidebar-scrollbox', MDBookSidebarScrollbox);


// ---------------------------------------------------------------------------
// Support for dynamically adding headers to the sidebar.

(function() {
    // This is used to detect which direction the page has scrolled since the
    // last scroll event.
    let lastKnownScrollPosition = 0;
    // This is the threshold in px from the top of the screen where it will
    // consider a header the "current" header when scrolling down.
    const defaultDownThreshold = 150;
    // Same as defaultDownThreshold, except when scrolling up.
    const defaultUpThreshold = 300;
    // The threshold is a virtual horizontal line on the screen where it
    // considers the "current" header to be above the line. The threshold is
    // modified dynamically to handle headers that are near the bottom of the
    // screen, and to slightly offset the behavior when scrolling up vs down.
    let threshold = defaultDownThreshold;
    // This is used to disable updates while scrolling. This is needed when
    // clicking the header in the sidebar, which triggers a scroll event. It
    // is somewhat finicky to detect when the scroll has finished, so this
    // uses a relatively dumb system of disabling scroll updates for a short
    // time after the click.
    let disableScroll = false;
    // Array of header elements on the page.
    let headers;
    // Array of li elements that are initially collapsed headers in the sidebar.
    // I'm not sure why eslint seems to have a false positive here.
    // eslint-disable-next-line prefer-const
    let headerToggles = [];
    // This is a debugging tool for the threshold which you can enable in the console.
    let thresholdDebug = false;

    // Updates the threshold based on the scroll position.
    function updateThreshold() {
        const scrollTop = window.pageYOffset || document.documentElement.scrollTop;
        const windowHeight = window.innerHeight;
        const documentHeight = document.documentElement.scrollHeight;

        // The number of pixels below the viewport, at most documentHeight.
        // This is used to push the threshold down to the bottom of the page
        // as the user scrolls towards the bottom.
        const pixelsBelow = Math.max(0, documentHeight - (scrollTop + windowHeight));
        // The number of pixels above the viewport, at least defaultDownThreshold.
        // Similar to pixelsBelow, this is used to push the threshold back towards
        // the top when reaching the top of the page.
        const pixelsAbove = Math.max(0, defaultDownThreshold - scrollTop);
        // How much the threshold should be offset once it gets close to the
        // bottom of the page.
        const bottomAdd = Math.max(0, windowHeight - pixelsBelow - defaultDownThreshold);
        let adjustedBottomAdd = bottomAdd;

        // Adjusts bottomAdd for a small document. The calculation above
        // assumes the document is at least twice the windowheight in size. If
        // it is less than that, then bottomAdd needs to be shrunk
        // proportional to the difference in size.
        if (documentHeight < windowHeight * 2) {
            const maxPixelsBelow = documentHeight - windowHeight;
            const t = 1 - pixelsBelow / Math.max(1, maxPixelsBelow);
            const clamp = Math.max(0, Math.min(1, t));
            adjustedBottomAdd *= clamp;
        }

        let scrollingDown = true;
        if (scrollTop < lastKnownScrollPosition) {
            scrollingDown = false;
        }

        if (scrollingDown) {
            // When scrolling down, move the threshold up towards the default
            // downwards threshold position. If near the bottom of the page,
            // adjustedBottomAdd will offset the threshold towards the bottom
            // of the page.
            const amountScrolledDown = scrollTop - lastKnownScrollPosition;
            const adjustedDefault = defaultDownThreshold + adjustedBottomAdd;
            threshold = Math.max(adjustedDefault, threshold - amountScrolledDown);
        } else {
            // When scrolling up, move the threshold down towards the default
            // upwards threshold position. If near the bottom of the page,
            // quickly transition the threshold back up where it normally
            // belongs.
            const amountScrolledUp = lastKnownScrollPosition - scrollTop;
            const adjustedDefault = defaultUpThreshold - pixelsAbove
                + Math.max(0, adjustedBottomAdd - defaultDownThreshold);
            threshold = Math.min(adjustedDefault, threshold + amountScrolledUp);
        }

        if (documentHeight <= windowHeight) {
            threshold = 0;
        }

        if (thresholdDebug) {
            const id = 'mdbook-threshold-debug-data';
            let data = document.getElementById(id);
            if (data === null) {
                data = document.createElement('div');
                data.id = id;
                data.style.cssText = `
                    position: fixed;
                    top: 50px;
                    right: 10px;
                    background-color: 0xeeeeee;
                    z-index: 9999;
                    pointer-events: none;
                `;
                document.body.appendChild(data);
            }
            data.innerHTML = `
                <table>
                  <tr><td>documentHeight</td><td>${documentHeight.toFixed(1)}</td></tr>
                  <tr><td>windowHeight</td><td>${windowHeight.toFixed(1)}</td></tr>
                  <tr><td>scrollTop</td><td>${scrollTop.toFixed(1)}</td></tr>
                  <tr><td>pixelsAbove</td><td>${pixelsAbove.toFixed(1)}</td></tr>
                  <tr><td>pixelsBelow</td><td>${pixelsBelow.toFixed(1)}</td></tr>
                  <tr><td>bottomAdd</td><td>${bottomAdd.toFixed(1)}</td></tr>
                  <tr><td>adjustedBottomAdd</td><td>${adjustedBottomAdd.toFixed(1)}</td></tr>
                  <tr><td>scrollingDown</td><td>${scrollingDown}</td></tr>
                  <tr><td>threshold</td><td>${threshold.toFixed(1)}</td></tr>
                </table>
            `;
            drawDebugLine();
        }

        lastKnownScrollPosition = scrollTop;
    }

    function drawDebugLine() {
        if (!document.body) {
            return;
        }
        const id = 'mdbook-threshold-debug-line';
        const existingLine = document.getElementById(id);
        if (existingLine) {
            existingLine.remove();
        }
        const line = document.createElement('div');
        line.id = id;
        line.style.cssText = `
            position: fixed;
            top: ${threshold}px;
            left: 0;
            width: 100vw;
            height: 2px;
            background-color: red;
            z-index: 9999;
            pointer-events: none;
        `;
        document.body.appendChild(line);
    }

    function mdbookEnableThresholdDebug() {
        thresholdDebug = true;
        updateThreshold();
        drawDebugLine();
    }

    window.mdbookEnableThresholdDebug = mdbookEnableThresholdDebug;

    // Updates which headers in the sidebar should be expanded. If the current
    // header is inside a collapsed group, then it, and all its parents should
    // be expanded.
    function updateHeaderExpanded(currentA) {
        // Add expanded to all header-item li ancestors.
        let current = currentA.parentElement;
        while (current) {
            if (current.tagName === 'LI' && current.classList.contains('header-item')) {
                current.classList.add('expanded');
            }
            current = current.parentElement;
        }
    }

    // Updates which header is marked as the "current" header in the sidebar.
    // This is done with a virtual Y threshold, where headers at or below
    // that line will be considered the current one.
    function updateCurrentHeader() {
        if (!headers || !headers.length) {
            return;
        }

        // Reset the classes, which will be rebuilt below.
        const els = document.getElementsByClassName('current-header');
        for (const el of els) {
            el.classList.remove('current-header');
        }
        for (const toggle of headerToggles) {
            toggle.classList.remove('expanded');
        }

        // Find the last header that is above the threshold.
        let lastHeader = null;
        for (const header of headers) {
            const rect = header.getBoundingClientRect();
            if (rect.top <= threshold) {
                lastHeader = header;
            } else {
                break;
            }
        }
        if (lastHeader === null) {
            lastHeader = headers[0];
            const rect = lastHeader.getBoundingClientRect();
            const windowHeight = window.innerHeight;
            if (rect.top >= windowHeight) {
                return;
            }
        }

        // Get the anchor in the summary.
        const href = '#' + lastHeader.id;
        const a = [...document.querySelectorAll('.header-in-summary')]
            .find(element => element.getAttribute('href') === href);
        if (!a) {
            return;
        }

        a.classList.add('current-header');

        updateHeaderExpanded(a);
    }

    // Updates which header is "current" based on the threshold line.
    function reloadCurrentHeader() {
        if (disableScroll) {
            return;
        }
        updateThreshold();
        updateCurrentHeader();
    }


    // When clicking on a header in the sidebar, this adjusts the threshold so
    // that it is located next to the header. This is so that header becomes
    // "current".
    function headerThresholdClick(event) {
        // See disableScroll description why this is done.
        disableScroll = true;
        setTimeout(() => {
            disableScroll = false;
        }, 100);
        // requestAnimationFrame is used to delay the update of the "current"
        // header until after the scroll is done, and the header is in the new
        // position.
        requestAnimationFrame(() => {
            requestAnimationFrame(() => {
                // Closest is needed because if it has child elements like <code>.
                const a = event.target.closest('a');
                const href = a.getAttribute('href');
                const targetId = href.substring(1);
                const targetElement = document.getElementById(targetId);
                if (targetElement) {
                    threshold = targetElement.getBoundingClientRect().bottom;
                    updateCurrentHeader();
                }
            });
        });
    }

    // Takes the nodes from the given head and copies them over to the
    // destination, along with some filtering.
    function filterHeader(source, dest) {
        const clone = source.cloneNode(true);
        clone.querySelectorAll('mark').forEach(mark => {
            mark.replaceWith(...mark.childNodes);
        });
        dest.append(...clone.childNodes);
    }

    // Scans page for headers and adds them to the sidebar.
    document.addEventListener('DOMContentLoaded', function() {
        const activeSection = document.querySelector('#mdbook-sidebar .active');
        if (activeSection === null) {
            return;
        }

        const main = document.getElementsByTagName('main')[0];
        headers = Array.from(main.querySelectorAll('h2, h3, h4, h5, h6'))
            .filter(h => h.id !== '' && h.children.length && h.children[0].tagName === 'A');

        if (headers.length === 0) {
            return;
        }

        // Build a tree of headers in the sidebar.

        const stack = [];

        const firstLevel = parseInt(headers[0].tagName.charAt(1));
        for (let i = 1; i < firstLevel; i++) {
            const ol = document.createElement('ol');
            ol.classList.add('section');
            if (stack.length > 0) {
                stack[stack.length - 1].ol.appendChild(ol);
            }
            stack.push({level: i + 1, ol: ol});
        }

        // The level where it will start folding deeply nested headers.
        const foldLevel = 3;

        for (let i = 0; i < headers.length; i++) {
            const header = headers[i];
            const level = parseInt(header.tagName.charAt(1));

            const currentLevel = stack[stack.length - 1].level;
            if (level > currentLevel) {
                // Begin nesting to this level.
                for (let nextLevel = currentLevel + 1; nextLevel <= level; nextLevel++) {
                    const ol = document.createElement('ol');
                    ol.classList.add('section');
                    const last = stack[stack.length - 1];
                    const lastChild = last.ol.lastChild;
                    // Handle the case where jumping more than one nesting
                    // level, which doesn't have a list item to place this new
                    // list inside of.
                    if (lastChild) {
                        lastChild.appendChild(ol);
                    } else {
                        last.ol.appendChild(ol);
                    }
                    stack.push({level: nextLevel, ol: ol});
                }
            } else if (level < currentLevel) {
                while (stack.length > 1 && stack[stack.length - 1].level > level) {
                    stack.pop();
                }
            }

            const li = document.createElement('li');
            li.classList.add('header-item');
            li.classList.add('expanded');
            if (level < foldLevel) {
                li.classList.add('expanded');
            }
            const span = document.createElement('span');
            span.classList.add('chapter-link-wrapper');
            const a = document.createElement('a');
            span.appendChild(a);
            a.href = '#' + header.id;
            a.classList.add('header-in-summary');
            filterHeader(header.children[0], a);
            a.addEventListener('click', headerThresholdClick);
            const nextHeader = headers[i + 1];
            if (nextHeader !== undefined) {
                const nextLevel = parseInt(nextHeader.tagName.charAt(1));
                if (nextLevel > level && level >= foldLevel) {
                    const toggle = document.createElement('a');
                    toggle.classList.add('chapter-fold-toggle');
                    toggle.classList.add('header-toggle');
                    toggle.addEventListener('click', () => {
                        li.classList.toggle('expanded');
                    });
                    const toggleDiv = document.createElement('div');
                    toggleDiv.textContent = '❱';
                    toggle.appendChild(toggleDiv);
                    span.appendChild(toggle);
                    headerToggles.push(li);
                }
            }
            li.appendChild(span);

            const currentParent = stack[stack.length - 1];
            currentParent.ol.appendChild(li);
        }

        const onThisPage = document.createElement('div');
        onThisPage.classList.add('on-this-page');
        onThisPage.append(stack[0].ol);
        const activeItemSpan = activeSection.parentElement;
        activeItemSpan.after(onThisPage);
    });

    document.addEventListener('DOMContentLoaded', reloadCurrentHeader);
    document.addEventListener('scroll', reloadCurrentHeader, { passive: true });
})();

