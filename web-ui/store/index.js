export const mutations = {
  merge (state, array) {
    let key  = array[0];
    let data = array[1];
    state[key] = { ...state[key], ...data }
  },
};

export const state = () => ({
  auth: {
    token:          null,
    id:             null,
    email:          null,
    nickname:       null,
    states:         null,
    avatar:         null,
    profile:        null,
    urlName:        null,
    noteCount:      null,
    followerCount:  null,
    followingCount: null,
  },

  snackbar: {
    message: '',
  },

  color: {

    // Bar
    barStart:                   '#f5f5f5',
    barEnd:                     '#eee',

    // Breadcrummbs
    breadcrumbsFore:            '#757575',
    breadcrumbsForeHover:       '#616161',
    breadcrumbsForeActive:      '#4ab7bd',
    breadcrumbsForeActiveHover: '#00ACC1',
    breadcrumbsDelimFore:       '#bdbdbd',

    // Btn
    btnBack:                    '#4ab7bd',
    btnSimpleBack:              '#9e9e9e',
    btnImportantBack:           '#e91e63',
    btnDangerBack:              '#ff5251',
    btnErrorBack:               '#ff5251',
    btnFore:                    '#fff',
    btnSimpleFore:              '#fff',
    btnImportantFore:           '#fff',
    btnDangerFore:              '#fff',
    btnErrorFore:               '#fff',

    // Card
    cardTitleBack:              '#4ab7bd',
    cardTitleFore:              '#fff',
    cardTextBack:               '#fff',
    cardTextFore:               '#616161',
    cardActionsBack:            '#fafafa',
    cardBorder:                 '#eee',

    // Chat
    chatGuestBack:              '#bdbdbd',
    chatGuestFore:              '#fff',
    chatHostBack:               '#f5f5f5',
    chatHostFore:               '#616161',
    chatSystemBack:             '#eee',
    chatSystemFore:             '#616161',

    // Chart
    chartFore:                  '#757575',
    chartBacks:                 ['rgba(96, 125, 139, .2)', 'rgba( 0, 188, 212, .1)'],
    chartBacksActive:           ['rgba(96, 125, 139, .4)', 'rgba( 0, 188, 212, .2)'],
    chartToolbarBack:           'rgba(98, 125, 137, .9)',

    // Checkbox
    checkboxesFore:             '#757575',
    checkboxesForeActive:       '#616161',
    checkboxesIcon:             '#9e9e9e',
    checkboxesIconActive:       '#757575',

    // Chip
    chipBack:                   '#bdbdbd',
    chipFore:                   '#fff',
    chipLinkBack:               '#4ab7bd',
    chipLinkFore:               '#fff',

    // Disabled
    disabledFore:               '#bdbdbd',
    disabledBack:               '#e0e0e0',
    disabledIcon:               '#bdbdbd',
    disabledBorder:             '#e0e0e0',

    // Error
    errorFore:                  '#ff5251',
    errorBack:                  '#ff5251',
    errorIcon:                  '#ff5251',
    errorBorder:                '#ff5251',

    // Footer
    footerBack:                 '#f5f5f5',
    footerFore:                 '#616161',
    footerBorder:               '#eee',

    // Link
    linkFore:                   '#4ab7bd',
    linkForeHover:              '#0097a7',

    // Navigation
    navBack:                    '#fafafa',
    navBackActive:              '#f5f5f5',
    navTitleFore:               '#9e9e9e',
    navTitleForeActive:         '#757575',
    navIconFore:                '#9e9e9e',
    navIconForeActive:          '#757575',
    navBorder:                  '#eee',

    // Page
    pageBack:                   '#fff',
    pageFore:                   '#616161',

    // Pagination
    paginationBack:             '#fff',
    paginationBackActive:       '#4ab7bd',
    paginationFore:             '#757575',
    paginationForeActive:       '#fff',
    paginationIcon:             '#9e9e9e',
    paginationMoreFore:         '#757575',
    paginationTotalFore:        '#9e9e9e',
    paginationBorder:           '#eee',
    paginationBorderActive:     '#4ab7bd',

    // Progress
    progress:                   '#4ab7bd',

    // Radios
    radiosFore:                 '#757575',
    radiosForeActive:           '#616161',
    radiosIcon:                 '#9e9e9e',
    radiosIconActive:           '#757575',

    // Scroll
    scrollBack:                 '#4ab7bd',
    scrollFore:                 '#fff',

    // Search
    searchForeFocused:          '#4ab7bd',
    searchIcon:                 '#bdbdbd',
    searchIconFocused:          '#4ab7bd',
    searchBorder:               '#e0e0e0',
    searchBorderFocused:        '#4ab7bd',
    searchChipBack:             '#bdbdbd',
    searchChipBackActive:       '#4ab7bd',
    searchChipBackFocused:      '#4ab7bd',
    searchChipFore:             '#fff',
    searchChipForeActive:       '#fff',
    searchChipForeFocused:      '#fff',
    searchChipTitle:            '#9e9e9e',
    searchChipDelim:            '#eee',

    // Select
    selectFore:                 '#757575',
    selectForeActive:           '#616161',
    selectForeFocused:          '#4ab7bd',
    selectIcon:                 '#9e9e9e',
    selectIconFocused:          '#4ab7bd',
    selectIconActive:           '#757575',
    selectIconPrev:             '#9e9e9e',
    selectIconNext:             '#9e9e9e',
    selectBorder:               '#e0e0e0',
    selectBorderFocused:        '#4ab7bd',
    selectBack:                 '#fff',
    selectBackActive:           '#eee',
    selectBackHover:            '#f5f5f5',

    // SnackBar
    snackbarBack:               '#4ab7bd',
    snackbarFore:               '#fff',

    // Switch
    switchFore:                 '#757575',
    switchFalseFore:            '#757575',
    switchTrueFore:             '#616161',
    switchFalseBack:            '#e0e0e0',
    switchTrueBack:             '#9e9e9e',

    // Table
    tableBorder:                '#eee',
    tdHeaderBack:               '#4ab7bd',
    tdHeaderFore:               '#fff',
    tdHeaderBorder:             '#fff',
    tdBack:                     '#fff',
    tdBackCur:                  '#fce4ec',
    tdBackSel:                  '#e0f7fa',
    tdBackHover:                '#f5f5f5',
    tdFore:                     '#616161',
    tdBorder:                   '#eee',
    tdLinkBackHover:            '#fafafa',
    tdRateBack:                 '#4dd0e1',

    // TableMany
    tableManyFore:              '#757575',
    tableManyForeActive:        '#616161',
    tableManyForeFocused:       '#4ab7bd',
    tableManyBackActive:        '#eee',
    tableManyBackHover:         '#f5f5f5',

    // TableOne
    tableOneFore:               '#757575',
    tableOneForeActive:         '#616161',
    tableOneForeFocused:        '#4ab7bd',
    tableOneBackActive:         '#eee',
    tableOneBackHover:          '#f5f5f5',

    // Textarea
    textareaFore:               '#757575',
    textareaForeActive:         '#616161',
    textareaForeFocused:        '#4ab7bd',
    textareaBorder:             '#e0e0e0',
    textareaBorderFocused:      '#4ab7bd',
    textareaIcon:               '#9e9e9e',
    textareaIconFocused:        '#4ab7bd',

    // Text field
    textFieldFore:              '#757575',
    textFieldForeActive:        '#616161',
    textFieldForeFocused:       '#4ab7bd',
    textFieldBorder:            '#e0e0e0',
    textFieldBorderFocused:     '#4ab7bd',
    textFieldIcon:              '#9e9e9e',
    textFieldIconFocused:       '#4ab7bd',
    textFieldAppend:            '#757575',

    // Toolbar
    toolbarBack:                '#4ab7bd',
  },
});
