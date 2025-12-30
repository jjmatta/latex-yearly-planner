{{- $today := .Body.Day -}}

\begin{minipage}[t]{\myLenTriCol}
\vspace{.5\baselineskip}
{{template "schedule.tpl" dict "Cfg" .Cfg "Day" .Body.Day}}
  \vspace{\dimexpr2mm+.3pt}

{{- if .Cfg.CalAfterSchedule -}}
{{- template "monthTabularV2.tpl" dict "Month" .Body.Month "Today" $today "Cfg" .Cfg -}}
{{- end -}}
\end{minipage}%
\hspace{\myLenTriColSep}%
\begin{minipage}[t]{\dimexpr2\myLenTriCol+\myLenTriColSep}
  \vspace{.5\baselineskip}
  \myUnderline{Notes $\vert$ {{ $today.LinkLeaf "More" "Additional notes" }}}
  \vbox to \dimexpr\textheight-\pagetotal-\myLenLineHeightButLine\relax {%
    \leaders\hbox to \linewidth{\textcolor{\myColorGray}{\rule{0pt}{\myLenLineHeightButLine}\hrulefill}}\vfil
  }
\end{minipage}
\par\pagebreak
