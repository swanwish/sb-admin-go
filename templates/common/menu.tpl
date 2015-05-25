{{define "menu"}}
<ul class="nav" id="side-menu">
    <li class="sidebar-search">
        <div class="input-group custom-search-form">
            <input type="text" class="form-control" placeholder="Search...">
                                <span class="input-group-btn">
                                <button class="btn btn-default" type="button">
                                    <i class="fa fa-search"></i>
                                </button>
                            </span>
        </div>
        <!-- /input-group -->
    </li>
    <li>
        <a href="/view/index"><i class="fa fa-dashboard fa-fw"></i> Dashboard</a>
    </li>
    <li>
        <a href="#"><i class="fa fa-bar-chart-o fa-fw"></i> Charts<span class="fa arrow"></span></a>
        <ul class="nav nav-second-level">
            <li>
                <a href="/view/flot">Flot Charts</a>
            </li>
            <li>
                <a href="/view/morris">Morris.js Charts</a>
            </li>
        </ul>
        <!-- /.nav-second-level -->
    </li>
    <li>
        <a href="/view/tables"><i class="fa fa-table fa-fw"></i> Tables</a>
    </li>
    <li>
        <a href="/view/forms"><i class="fa fa-edit fa-fw"></i> Forms</a>
    </li>
    <li>
        <a href="#"><i class="fa fa-wrench fa-fw"></i> UI Elements<span class="fa arrow"></span></a>
        <ul class="nav nav-second-level">
            <li>
                <a href="/view/panels-wells">Panels and Wells</a>
            </li>
            <li>
                <a href="/view/buttons">Buttons</a>
            </li>
            <li>
                <a href="/view/notifications">Notifications</a>
            </li>
            <li>
                <a href="/view/typography">Typography</a>
            </li>
            <li>
                <a href="/view/icons"> Icons</a>
            </li>
            <li>
                <a href="/view/grid">Grid</a>
            </li>
        </ul>
        <!-- /.nav-second-level -->
    </li>
    <li>
        <a href="#"><i class="fa fa-sitemap fa-fw"></i> Multi-Level Dropdown<span class="fa arrow"></span></a>
        <ul class="nav nav-second-level">
            <li>
                <a href="#">Second Level Item</a>
            </li>
            <li>
                <a href="#">Second Level Item</a>
            </li>
            <li>
                <a href="#">Third Level <span class="fa arrow"></span></a>
                <ul class="nav nav-third-level">
                    <li>
                        <a href="#">Third Level Item</a>
                    </li>
                    <li>
                        <a href="#">Third Level Item</a>
                    </li>
                    <li>
                        <a href="#">Third Level Item</a>
                    </li>
                    <li>
                        <a href="#">Third Level Item</a>
                    </li>
                </ul>
                <!-- /.nav-third-level -->
            </li>
        </ul>
        <!-- /.nav-second-level -->
    </li>
    <li>
        <a href="#"><i class="fa fa-files-o fa-fw"></i> Sample Pages<span class="fa arrow"></span></a>
        <ul class="nav nav-second-level">
            <li>
                <a class="active" href="/view/blank">Blank Page</a>
            </li>
            <li>
                <a href="/view/login">Login Page</a>
            </li>
        </ul>
        <!-- /.nav-second-level -->
    </li>
</ul>
{{end}}