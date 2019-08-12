{{define "content"}}
	<main>
            <div>
                <!-- Image of patient should be a burger that opens Account, Edit-Profile, and LogOut -->
                <a href="drProfile.html"> <img src="ptImages/ptProfile.png" title="PTsPictue" width="55" height="50" /></a>
                <a href="account.html">Account</a>
                <a href="drApplication.html">Update Profile</a>
                <a href="index.html">Log Out</a>
            </div>
            <a href="findPt.html">Find Patients</a>
            <a href="drMessages.html">Messages</a>
            <a href="drWallet.html">Wallet</a>
            <a href="help.html">Help</a>
            <a href="ptsBooked.html">Manage Patients</a>
        <hr/>
        
 <h2>YOU CURRENTLY HAVE NO APPOINTMENTS BOOKED</h2>
 <p>If you have agreed to see a patient but do not see them on this page, please refer to the <a href="help.html">Help</a> section.</a>
	</main>
{{end}}