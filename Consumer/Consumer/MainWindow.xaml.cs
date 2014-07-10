using System;
using System.ComponentModel;
using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Threading;
using System.Windows;

namespace Consumer
{
    /// <summary>
    /// Interaktionslogik für MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {

        WebClient client;
        String server, newClient;

        public MainWindow()
        {
            InitializeComponent();

            new Thread(() =>
            {
                Thread.CurrentThread.IsBackground = true;

                int recv;
                byte[] data = new byte[1024];
                IPEndPoint ipep = new IPEndPoint(IPAddress.Any, 12345);

                Socket newsock = new Socket(AddressFamily.InterNetwork,
                                SocketType.Dgram, ProtocolType.Udp);

                newsock.Bind(ipep);

                IPEndPoint sender = new IPEndPoint(IPAddress.Any, 0);
                EndPoint Remote = (EndPoint)(sender);


                while (true)
                {
                    data = new byte[1024];
                    recv = newsock.ReceiveFrom(data, ref Remote);

                    newClient = Encoding.UTF8.GetString(data, 0, recv);

                    lstServers.Dispatcher.BeginInvoke((Action) (() =>
                    {
                        lstServers.Items.Add(newClient + ":12345");
                        newClient = "";
                    }));

                    newsock.SendTo(data, recv, SocketFlags.None, Remote);
                }

            }).Start();


        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {
            executeRequest(txtCommand.Text);
            executeRequest("_running");
        }

        private void executeRequest(String command)
        {
            if (lstServers.SelectedIndex == -1)
            {
                return;
            }

            var worker = new BackgroundWorker();

            server = lstServers.SelectedItem.ToString();

            worker.DoWork += (sender, args) => {
                try
                {
                    args.Result = new WebClient().DownloadString("http://" + server + "/" + Uri.EscapeDataString(command));
                }
                catch (Exception e)
                {
                    args.Result = "Fehler bei der Verbindung.";
                }
            };

            worker.RunWorkerCompleted += (sender, e) => {
                if (e.Error != null) {
                    output.Text = "Error: " + e.Error.Message;
                } else {
                    output.Text = e.Result.ToString();
                }
            };

            worker.RunWorkerAsync();

        }

        private void Button_Click_1(object sender, RoutedEventArgs e)
        {
            lstServers.Items.Add(txtServer.Text);
        }

        private void Button_Click_3(object sender, RoutedEventArgs e)
        {
            executeRequest("cmd /c echo list disk | diskpart");
        }

        private void Button_Click_4(object sender, RoutedEventArgs e)
        {
            executeRequest("");
        }

        private void Button_Click_5(object sender, RoutedEventArgs e)
        {
            executeRequest("wpeutil reboot");
        }

        private void Button_Click_2(object sender, RoutedEventArgs e)
        {
            executeRequest("_info");
        }

        private void Button_Click_7(object sender, RoutedEventArgs e)
        {
            executeRequest("wmic computersystem get Caption, Manufacturer, Model");
        }

        private void Button_Click_8(object sender, RoutedEventArgs e)
        {
            executeRequest("_running");
        }
    }
}
